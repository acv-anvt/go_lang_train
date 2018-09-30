Sidekiq.configure_client do |config|
  config.redis = { db: 0 }
end

Sidekiq.configure_server do |config|
  config.redis = { db: 0 }
end
p time_start = Time.current
p start = time_start.to_i
30000.times do |e|
  Sidekiq::Client.push("queue" => "go_queue","class" => "GoWorker","args"  => ["I am Running #{e}"])
end
p end_time = Time.current.to_i
p "Total: #{end_time-start} (s) => #{(end_time-start)*1000*1000} (μs)"
"Total: 5 (s) => 5000000 (μs)"

23:35:50
23:36:03
= 13 (s)

"Total: 11 (s) => 11000000 (μs)"
23:52:02
23:52:13
= 11 (s)

"Total: 12 (s) => 12000000 (μs)"
23:53:26
23:53:38
=12 (s)

p time_start = Time.current
p start = time_start.to_i
30000.times do |e|
  PrintTestJob.perform_later "I am Running #{e}"
end
p end_time = Time.current.to_i
p "Total: #{end_time-start} (s) => #{(end_time-start)*1000*1000} (μs)"
23:41:47
"Total: 14 (s) => 14000000 (μs)"

16:41:47
16:42:41
= 54 (s)

"Total: 20 (s) => 20000000 (μs)"
23:55:24
23:56:11
= 47 (s)


"Total: 14 (s) => 14000000 (μs)"
23:57:59
23:58:42
=43 (s)
