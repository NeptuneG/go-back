# frozen_string_literal: true

class ScrapeBillboardLiveTokyoJob
  include Faktory::Job

  def perform(year_month)
    live_events = ScrapeBillboardLiveTokyo.new(year_month).call
    Redis.new(host: 'redis').publish('live_events', live_events.to_json)
    logger.info("#{live_events.count} events published")
  rescue StandardError => e
    logger.error(e)
    raise e
  end

  private

  def logger
    @logger ||= Logger.new(STDOUT)
  end
end
