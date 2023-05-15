require 'sinatra'
require 'uri'
require 'net/http'
require 'opentelemetry/sdk'
require 'opentelemetry/exporter/otlp'
require 'opentelemetry/instrumentation/all'

# this block is the only code that needs to be added to the application
# https://opentelemetry.io/docs/instrumentation/ruby/automatic/
OpenTelemetry::SDK.configure do |c|
  c.service_name = 'uppercase'
  c.use_all() # enables all instrumentation!
end

# continue original application code
get '/uppercase' do
  subject = params['subject']
  finalupper_addr = ENV["FINALUPPER_ADDR"]
  uri = URI("#{finalupper_addr}/finalupper?subject=#{subject}")
  
  res = Net::HTTP.get_response(uri)
  res.body
end
