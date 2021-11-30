# frozen_string_literal: true

class ScrapeBillboardLiveTokyoJob
  include Faktory::Job

  def perform(year_month)
    ScrapeBillboardLiveTokyo.new(year_month).call
  rescue StandardError => e
    Logger.new(STDOUT).error(e)
    raise e
  end
end
