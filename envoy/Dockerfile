FROM  envoyproxy/envoy:v1.12.2
WORKDIR /etc/envoy
COPY ./envoy.yaml .
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
EXPOSE 80