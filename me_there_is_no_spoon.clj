(ns Player
  (:gen-class))

;; Don't let the machines win. You are humanity's last hope...

(defn output [msg] (println msg) (flush))
(defn debug [msg] (binding [*out* *err*] (println msg) (flush)))
(defn format-node [x]
  (str (x 0) " " (x 1)))
(defn doesNodeExist [lon X Y]
  (if (some #(= [X Y] %) lon)
    (do [X Y])
    (do [-1 -1]))
  )
(defn getDoValue [lon x]
  (let [node
        (some
         (fn [a] (when (and
                        (= (x 0) (a 0))
                        (< (x 1) (a 1)) )
                   a ) )
         lon)]
    (if (not node)
      (do [-1 -1])
      (do node))
    )
  )
(defn getRiValue [lon x]
  (let [node
        (some
         (fn [a] (when (and
                        (< (x 0) (a 0))
                        (= (x 1) (a 1)) )
                   a ) )
         lon)]
    (if (not node)
      (do [-1 -1])
      (do node))
    )
  )


(defn -main [& args]
  (let [width (read) height (read) _ (read-line)]
    (with-local-vars [list-of-in [] list-of-out '() list-of-node '[]]
      ;; width: the number of cells on the X axis
      ;; height: the number of cells on the Y axis
      (debug (str "# PERF INIT: " (.getTime (java.util.Date.)) ))
      (loop [i 0]
        (when (< i height)
          (let [line (into [] (clojure.string/split (read-line) #"" ))]
            ;; line: width characters, each either 0 or .
            ;; (debug (str i" -> " line))
            (var-set list-of-in (into (var-get list-of-in) line))
            (loop [j 0]
              (when (< j width)
                (when (= "0" (get line j))
                  (var-set list-of-node (conj (var-get list-of-node) [j i]))
                  )
              (recur (inc j))))
            (recur (inc i)))))

      (debug (str "# PERF PARSED: " (.getTime (java.util.Date.)) ))

      ;; (debug "Debug messages...")
      ;; (debug (var-get list-of-node))
      (var-set list-of-out
               (for [node (var-get list-of-node)]
                 (str (format-node node) " "
                      (format-node (getRiValue (var-get list-of-node) node)) " "
                      (format-node (getDoValue (var-get list-of-node) node))
                      )
                 ))
      ;; (doseq [out (var-get list-of-out)]
      ;;   (debug out))

      (debug (str "# PERF BUILDED: " (.getTime (java.util.Date.)) ))

      (doseq [out (var-get list-of-out)]
        (output out))
      (debug (str "# PERF ENDED: " (.getTime (java.util.Date.))))

      ;; Three coordinates: a node, its right neighbor, its bottom neighbor
      )))

(-main)
