FROM alpine

# https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.zip
ENV COMPACT_GOOGLE_CLOUD ./vendor/google-cloud-sdk.zip

RUN apk add --update python ca-certificates
# ============================================================================
ENV CLOUDSDK_PYTHON_SITEPACKAGES 1
# Install the Google Cloud SDK.
ADD $COMPACT_GOOGLE_CLOUD /tmp/google-cloud-sdk.zip
RUN unzip /tmp/google-cloud-sdk.zip -d / && rm /tmp/google-cloud-sdk.zip

# Add the Drone plugin
ADD drone-gke /bin/

ENTRYPOINT ["/bin/drone-gke"]
