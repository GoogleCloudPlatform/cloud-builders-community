$projectID = gcloud config get-value project;
docker build -t gcr.io/$projectID/docker-windows .;
echo "auth"
gcloud auth configure-docker --quiet;
if ($?) {
    docker push gcr.io/$projectID/docker-windows;
}
