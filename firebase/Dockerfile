FROM gcr.io/cloud-builders/npm

RUN npm i firebase-tools
COPY firebase.bash .

ENTRYPOINT ["/firebase.bash"]