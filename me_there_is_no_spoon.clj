(ns Player
  (:gen-class))

;; Don't let the machines win. You are humanity's last hope...

(defn output [msg] (println msg) (flush))
(defn debug [msg] (binding [*out* *err*] (println msg) (flush)))
(defn doesNodeExist [lon X Y]
  (if (some #(= [X Y] %) lon)
    (do [X Y])
    (do [-1 -1]))
  )

(defn -main [& args]
  (let [width (read) height (read) _ (read-line)]
    (with-local-vars [list-of-in [] list-of-out '() list-of-node '[]]
      ;; width: the number of cells on the X axis
      ;; height: the number of cells on the Y axis
      (loop [i 0]
        (when (< i height)
          (let [line (into [] (clojure.string/split (read-line) #"" ))]
            ;; line: width characters, each either 0 or .
            (debug (str i" -> " line))
            (var-set list-of-in (into (var-get list-of-in) line))
            (loop [j 0]
              (when (< j width)
                (when (= "0" (get line j))
                  (var-set list-of-node (conj (var-get list-of-node) [i j]))
                  )
              (recur (inc j))))
            (recur (inc i)))))

      ;; (debug "Debug messages...")
      (debug (var-get list-of-node))
      (var-set list-of-out
               (for [node (var-get list-of-node)]
                 (str (node 0) " "
                      (node 1) " "
                      ((doesNodeExist (var-get list-of-node) (+ 1 (get node 0)) (get node 1)) 0) " "
                      ((doesNodeExist (var-get list-of-node) (+ 1 (get node 0)) (get node 1)) 1) " "
                      ((doesNodeExist (var-get list-of-node) (get node 0) (+ 1 (get node 1))) 0) " "
                      ((doesNodeExist (var-get list-of-node) (get node 0) (+ 1 (get node 1))) 1)
                      )
                 ))
      (doseq [out (var-get list-of-out)]
        (debug out))

      (doseq [out (var-get list-of-out)]
        (output out))

      ;; Three coordinates: a node, its right neighbor, its bottom neighbor
      )))
