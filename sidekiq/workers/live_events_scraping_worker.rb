# frozen_string_literal: true

require_relative '../services/bbl_tokyo_scraper'

class LiveEventsScrapingWorker
  include Sidekiq::Worker

  sidekiq_options queue: :live_events, retry: false

  def perform(year_month)
    LiveEvents::BblTokyoScraper.new(year_month).call
  end
end
