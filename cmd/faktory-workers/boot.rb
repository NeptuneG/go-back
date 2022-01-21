# frozen_string_literal: true

require 'faktory'
require 'redis'
require 'yabeda/prometheus'

Dir[File.join('.', '**/*.rb')].sort.each { |file| require file }

Faktory.configure_worker do |_config|
  Yabeda::Prometheus::Exporter.start_metrics_server!
end
