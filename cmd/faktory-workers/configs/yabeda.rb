# frozen_string_literal: true

Yabeda.configure do
  group :live_events_scraper do
    counter :scraped_live_events_count, comment: 'Total number of scrapped live events', tags: %i[live_house count]
    # gauge     :whistles_active,  comment: 'Number of whistles ready to whistle'
    histogram :scrapping_runtime, buckets: [0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30, 60].freeze, tags: %i[live_house] do
      comment 'How long is taken to scrape live events'
      unit :seconds
    end
  end
end

Yabeda.configure!
