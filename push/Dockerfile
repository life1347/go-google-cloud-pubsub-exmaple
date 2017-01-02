FROM scratch
# ca-certificates.crt can be found under the /etc/ssl/certs/
# in any linux distribution 
ADD asset/ca-certificates.crt /etc/ssl/certs/
# You cann apply for free ssl cert at https://www.sslforfree.com/
ADD asset/certificate.crt /certificate.crt
ADD asset/private.key /private.key
ADD pushSubscription /pushSubscription
CMD ["/pushSubscription"]