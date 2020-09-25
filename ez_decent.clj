;; https://www.codingame.com/training/easy/the-descent
(ns Player
  (:gen-class))


(defn output [msg] (println msg) (flush))
(defn debug [msg] (binding [*out* *err*] (println msg) (flush)))

(defn -main [& args]
  (while true
    (def mountains {})
    (dotimes [iter 8]
      (let [mountainH (read)]
        (debug (str "mnt:" iter "Height:" mountainH))
        (def mountains (assoc mountains mountainH iter))
        )
      )

    (debug mountains)
    (def sorted-mountains (reverse (into (sorted-map) mountains)))
    (debug sorted-mountains)

    (output (last (first sorted-mountains)))
    (def sorted-mountains (pop sorted-mountains))
    ))
(-main)
