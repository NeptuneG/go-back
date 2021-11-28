# frozen_string_literal: true

class ScrapingLiveEventsWorker
  include Sidekiq::Worker

  sidekiq_options queue: :scraping_live_events, retry: false

  def perform(year_month)
    ScrapeBillboardLiveTokyo.new(year_month).call
  rescue StandardError => e
    Logger.new(STDOUT).error(e)
    raise e
  end
end
