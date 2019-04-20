#lang racket

(define url "friends.reckhard.ca")
(define envs '(("prod" 3000) ("dev" 5000) ("spike" 8000)))
(define servers '(("auth" 0) ("msg" 1)))

(define docker-traefik-info
  "version: \"3\"

services:
  reverse-proxy:
    image: traefik # The official Traefik docker image
    command: --docker # --api # Enables the web UI and tells Traefik to listen to docker
    ports:
      - \"80:80\"     # The HTTP port
      #- \"8080:8080\"     # The admin port
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # So that Traefik can listen to the Docker events
")




 ;; For making server instances
(define (make-server env svr)
  (format "  ~a:\n    image: ~a\n    env_file: ~a\n    environment:\n      - PORT=~a\n    labels:\n~a\n~a\n~a\n"
          (make-server-name (car env) (car svr))
          (make-server-image-name (car env) (car svr))
          (make-server-env-file (car env) (car svr))
          (+ (cadr env) (cadr svr))
          "      - traefik.enable=true"
          (make-server-frontend (car env) (car svr))
          (make-server-backend (cadr env) (cadr svr))))

(define (make-server-name env svr)
  (format "~a-~a-~a" svr "server" env))
(define (make-server-image-name env svr)
  (format "~a_~a_~a" env svr "server"))
(define (make-server-env-file env svr)
  (format "./~a.env" env))
(define (make-server-frontend env svr)
  (format "      - \"traefik.frontend.rule=Host:~a.~a.~a\"" svr env url))
(define (make-server-backend env svr)
  (format "      - \"traefik.port=~a\"" (+ svr env)))



;; For calling the make server instances
(define (make-env env)
  (for/list ([svr servers])
    (make-server env svr)))

(define (make-docker)
  (for/list ([e envs])
    (string-join (make-env e) "\n")))


;; Displaying the results
(display docker-traefik-info)
(display (string-join (make-docker) "\n"))
