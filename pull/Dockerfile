FROM scratch
# ca-certificates.crt can be found under the /etc/ssl/certs/
# in any linux distribution 
ADD asset/ca-certificates.crt /etc/ssl/certs/
# Please refer the section "How the Application Default Credentials work" at
# https://developers.google.com/identity/protocols/application-default-credentials
ADD asset/crt.json /crt.json
ADD pullSubscription /pullSubscription
CMD ["/pullSubscription"]