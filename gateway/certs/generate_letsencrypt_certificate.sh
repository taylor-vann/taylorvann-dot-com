# brian taylor vann
# taylovann.com

# generate_letsencrypt_certificate

email=brian.t.vann@gmail.com
domains=(
  taylorvann.com
  www.taylorvann.com
  statics.taylorvann.com
  www.statics.taylorvann.com
  authn.taylorvann.com
  www.authn.taylorvann.com
  mail.taylorvann.com
  www.mail.taylorvann.com
)

domain_string=""
function create_domain_string_list()
{
  for i in $@;
  do
    domain_string+=${i},; 
  done

  domain_string=${domain_string::-1}
}

create_domain_string_list ${domains[@]}

# echo ${email}
# echo ${domain_string}

certbot certonly --standalone --non-interactive --agree-tos --email ${email} --domains ${domain_string}