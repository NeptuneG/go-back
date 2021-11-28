# frozen_string_literal: true

require 'parallel'
require 'metainspector'

class ScrapeBillboardLiveTokyo
  def initialize(year_month)
    @year_month = year_month
    @logger = Logger.new(STDOUT)
  end

  def call
    scraped_live_events.each { |live_events| puts(live_events.to_json) }
  end

  private

  def scraped_live_events
    Parallel.map(live_event_urls, in_threads: 20) { |url| scrape_event(url) }.compact
  end

  def live_event_urls
    schedule_page.links.all.filter { |url| live_event_url?(url) }
  end

  def schedule_page
    MetaInspector.new(year_month_url)
  end

  def year_month_url
    home_page_url = 'http://www.billboard-live.com/pg/shop/show/index.php'
    "#{home_page_url}?date=#{@year_month}&mode=calendar&shop=1"
  end

  def live_event_url?(url)
    %r{http://www.billboard-live.com/pg/shop/show/index.php\?mode=\w+&event=\d+&shop=1$}.match?(url)
  end

  def scrape_event(live_event_url)
    @logger.info("Scraping #{live_event_url}")
    parsed_live_event_page = MetaInspector.new(live_event_url).parsed
    scrape_date = scrape_date(parsed_live_event_page)
    scrape_open_starts = scrape_open_starts(parsed_live_event_page)
    {
      live_house_name: 'Billboard Live TOKYO', url: live_event_url,
      title: scrape_title(parsed_live_event_page),
      description: scrape_description(parsed_live_event_page),
      price_info: scrape_price_info(parsed_live_event_page),
      stage_one_open_at: "#{scrape_date} #{scrape_open_starts[0]}",
      stage_one_start_at: "#{scrape_date} #{scrape_open_starts[1]}",
      stage_two_open_at: "#{scrape_date} #{scrape_open_starts[2]}",
      stage_two_start_at: "#{scrape_date} #{scrape_open_starts[3]}"
    }
  rescue StandardError => e
    @logger.error("Scraping #{live_event_url} failed: #{e.message}")
    {}
  end

  def scrape_title(parsed_live_event_page)
    trim(parsed_live_event_page.css('div.lf_area_liveinfo h3.lf_tokyo').first.text)
  end

  def scrape_description(parsed_live_event_page)
    trim(parsed_live_event_page.css('div.lf_area_liveinfo p')[2].text, keep_linefeed: true)
  end

  def scrape_price_info(parsed_live_event_page)
    trim(parsed_live_event_page.css('div.lf_txtarea_memberprice').first.parent.css('p').first.text)
  end

  def scrape_date(parsed_live_event_page)
    parsed_live_event_page.css('div.lf_area_liveinfo div.lf_ttl h3').text[...-3]
  end

  def scrape_open_starts(parsed_live_event_page)
    trim(parsed_live_event_page.css('div.lf_openstart').text).split.filter { |word| match_hour_format?(word) }
  end

  def trim(text, keep_linefeed: false)
    reg = keep_linefeed ? /[[:space:]&&[^\n]]+/ : /[[:space:]]+/
    text.strip.gsub(reg, ' ')
  end

  def match_hour_format?(timestamp)
    # HH:MM 24-hour with leading 0
    /^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$/.match?(timestamp)
  end
end
