FROM golang:alpine

# lest be shore where we working
WORKDIR /

RUN apk add --update python ca-certificates unzip git

# https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.zip
ENV COMPACT_GOOGLE_CLOUD ./vendor/google-cloud-sdk.zip

# Copy sdk in
ADD $COMPACT_GOOGLE_CLOUD /tmp/google-cloud-sdk.zip

# Unzip it
RUN unzip /tmp/google-cloud-sdk.zip -d / && rm /tmp/google-cloud-sdk.zip

# Copy shared gke install script
COPY vendor/install_gke /

# Run it
RUN ./install_gke

# Add gcloud to PATH
ENV PATH /google-cloud-sdk/bin:$PATH

# Cleanup
RUN apk del unzip

# =============================================================================

COPY . /go/src/github.com/yebo-ecommerce/drone-gke

WORKDIR /go/src/github.com/yebo-ecommerce/drone-gke

RUN go get -d -v

RUN go install -v
