Sidekiq.configure_client do |config|
  config.redis = { db: 0 }
end

Sidekiq.configure_server do |config|
  config.redis = { db: 0 }
end

p time_start = Time.current
p start = time_start.to_i
30000.times do |e|
  Sidekiq::Client.push("queue" => "go_queue","class" => "GoWorker","args"  => ["I am Inserting Person No. #{e}"])
end
p end_time = Time.current.to_i
p "Total: #{end_time-start} (s) => #{(end_time-start)*1000*1000} (μs)"

Enqueue: 22(s)

09:28:15
09:30:10
= 113 (s)


Enqueue: 21(s)
09:38:10
09:39:54
60+44 =104 (s)

p time_start = Time.current
p start = time_start.to_i
30000.times do |e|
  PrintTestJob.perform_later "I am Inserting Person No. #{e}"
end
p end_time = Time.current.to_i
p "Total: #{end_time-start} (s) => #{(end_time-start)*1000*1000} (μs)"

Enqueue: 22(s)
09:30:48
09:33:07
=139 (s)

Enqueue: 21(s)
09:41:37
09:43:44
=120+7 =127 (s)

09:54:24
09:56:23
