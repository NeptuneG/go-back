# frozen_string_literal: true

require 'faktory'
require 'redis'

Dir[File.join(".", "**/*.rb")].each { |file| require file }
