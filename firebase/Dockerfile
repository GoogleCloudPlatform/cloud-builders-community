FROM gcr.io/cloud-builders/npm

RUN npm i -g firebase-tools
COPY firebase.bash .

ENTRYPOINT ["/firebase.bash"]
