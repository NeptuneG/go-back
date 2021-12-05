# frozen_string_literal: true

class ScrapeBillboardLiveTokyoJob
  include Faktory::Job

  def perform(year_month)
    live_events = ScrapeBillboardLiveTokyo.new(year_month).call
    # Redis.new(host: 'redis').publish('live_events', live_events.to_json)
    redis_client.lpush('screped_live_events', live_events.map(&:to_json))
    logger.info("#{live_events.count} events scraped for #{year_month}")
  rescue StandardError => e
    logger.error(e)
    raise e
  end

  private

  def logger
    @logger ||= Logger.new(STDOUT)
  end

  def redis_client
    @redis_client ||= Redis.new(host: 'redis')
  end
end
