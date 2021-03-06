# frozen_string_literal: true

class ScrapeBillboardLiveTokyoJob
  include Faktory::Job

  def perform(year_month)
    Yabeda.live_events_scraper.scrapping_runtime.measure(live_house: :billboard_live_tokyo) do
      @live_events = ScrapeBillboardLiveTokyo.new(year_month).call
    end
    redis_client.lpush('screped_live_events', @live_events.map(&:to_json))
    Yabeda.live_events_scraper.scraped_live_events_count.increment(
      { live_house: :billboard_live_tokyo, count: @live_events.count },
      by: 1
    )
    logger.info("#{@live_events.count} events scraped for #{year_month}")
  end

  private

  def logger
    @logger ||= Logger.new($stdout)
  end

  def redis_client
    @redis_client ||= Redis.new(host: ENV['REDIS_MQ_SERVICE_HOST'], port: ENV['REDIS_MQ_SERVICE_PORT'])
  end
end
