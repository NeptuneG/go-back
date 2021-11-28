# frozen_string_literal: true

require 'sidekiq'

Dir["#{File.dirname(__FILE__)}/workers/*.rb"].each { |file| require file }

redis_url = "redis://#{ENV.fetch('REDIS_HOST', 'redis')}:#{ENV.fetch('REDIS_PORT', 6379)}"

configuration = {
  url: redis_url,
  password: ENV['REDIS_PASSWORD'],
  namespace: 'sidekiq'
}

Sidekiq.configure_server do |config|
  config.redis = configuration
end

Sidekiq.configure_client do |config|
  config.redis = configuration
end
