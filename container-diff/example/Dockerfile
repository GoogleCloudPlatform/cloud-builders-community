FROM launcher.gcr.io/google/ubuntu16_04

# Generate a random 32-char string.
RUN cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1 | tee rand.txt
