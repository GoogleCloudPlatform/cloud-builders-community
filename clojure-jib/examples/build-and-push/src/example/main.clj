(ns example.main
  (:require [ring.adapter.jetty :refer [run-jetty]]))

(defn -main
  [& _]
  (run-jetty 
    (fn [_] {:status 200 :body "clojure-jib example"})
    {:port (Long/parseLong (System/getenv "PORT"))
     :join? false}))
