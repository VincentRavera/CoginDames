;; https://www.codingame.com/training/medium/shadows-of-the-knight-episode-1
(ns Player
  (:gen-class))

;; Auto-generated code below aims at helping you parse
;; the standard input according to the problem statement.

(defn output [msg] (println msg) (flush))
(defn debug [msg] (binding [*out* *err*] (println msg) (flush)))

(defn median [A B]
  (int (/ (+ A B) 2)))

(defn u [Y0 Yi YN]
  [Y0 (median Yi Y0) Yi])
(defn d [Y0 Yi YN]
  [Yi (median Yi YN) YN])

(defn r [X0 Xi XN]
  [Xi (median Xi XN) XN])
(defn l [X0 Xi XN]
  [X0 (median Xi X0) Xi])

(defn U [data]
  (assoc data :Y (apply u (:Y data))))

(defn D [data]
  (assoc data :Y (apply d (:Y data))))

(defn R [data]
  (assoc data :X (apply r (:X data))))

(defn L [data]
  (assoc data :X (apply l (:X data))))

(defn process [data bmbdir]
  (case bmbdir
    U (U data)
    D (D data)
    R (R data)
    L (L data)
    UR (U (R data))
    UL (U (L data))
    DR (D (R data))
    DL (D (L data))
    )
  )
(defn extract [data axis]
  (get (axis data) 1)
  )

(defn piss-output [data]
  (output (str (extract data :X) " " (extract data :Y))))

(defn -main [& args]
  (let [W (read) H (read) N (read) Xi (read) Yi (read) X0 0 Y0 0]
    (with-local-vars [data { :X [X0 Xi W] :Y [Y0 Yi H]}]
      ;; W: width of the building.
      ;; H: height of the building.
      ;; N: maximum number of turns before game over.
      ;; X0: Position sur W
      ;; Y0: Position sur H

      (while true
        (let [bombDir (read)]
          ;; bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
          (debug (str "Directions: " bombDir))
          (var-set data (process (var-get data) bombDir))
          (debug (var-get data))
          (piss-output (var-get data))

          ;; (debug "Debug messages...")

          ;; the location of the next window Batman should jump to.
          )))))
