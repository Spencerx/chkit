
##chkit

###Aliases:
  
###Usage  :
 Chkit is a terminal client for containerum.io powerful API
###Example:
  
###Flags  :
  -n, --namespace string   
  -p, --password string    account password
  -q, --quiet              quiet mode
  -u, --username string    account username

###Subcommands :
  + create : Create deployment or service
  + delete : Delete resource
  + doc : Print full chkit help
  + get : Get resource data
  + help [command] : Help about any command
  + login : Login to system
  + logs : View pod logs
  + rename : Rename resource
  + replace : Replace deployment or service
  + run : Run a solution
  + set : Set configuration variables
  + update : update chkit client
  + version : Print version
  

##version

###Aliases:
  
###Usage  :
 Print version
###Example:
  
###Flags  :

###Subcommands :
  

##update

###Aliases:
  
###Usage  :
 update chkit client
###Example:
  chkit update [from github|dir <path>] [--debug]
###Flags  :
      --debug   print debug information

###Subcommands :
  + from : 
  

##update from

###Aliases:
  
###Usage  :
 
###Example:
  
###Flags  :

###Subcommands :
  + dir : update from local directory
  + github : update from github releases
  

##from github

###Aliases:
  
###Usage  :
 update from github releases
###Example:
  
###Flags  :

###Subcommands :
  

##from dir

###Aliases:
  
###Usage  :
 update from local directory
###Example:
  chkit update from dir <path> [--debug]
###Flags  :

###Subcommands :
  

##set

###Aliases:
  
###Usage  :
 Set configuration variables
###Example:
  
###Flags  :

###Subcommands :
  + access : set namespace access
  + containerum-api : 
  + default-namespace : set default namespace
  + image : set container image in specific deployment
  + replicas : set deployment replicas
  

##set replicas

###Aliases:
  re, rep, repl, replica
###Usage  :
 Sets deployment replicas
###Example:
  chkit set replicas [-n namespace_label] [-d depl_label] [N_replicas]
###Flags  :
  -d, --deployment string   deployment name
  -r, --replicas uint       replicas, 1..15 (default 1)

###Subcommands :
  

##set image

###Aliases:
  imgs, img, im, images
###Usage  :
 Sets container image in specific deployment.
If deployment contains only one container, then uses that container by default.
###Example:
  
###Flags  :
  -c, --container string    container label
  -d, --deployment string   deployment label
  -f, --force               suppress confirmation
  -i, --image string        new image

###Subcommands :
  

##set default-namespace

###Aliases:
  def-ns, default-ns, defns, def-namespace
###Usage  :
 set default namespace
###Example:
  
###Flags  :

###Subcommands :
  

##set containerum-api

###Aliases:
  api, current-api, api-addr, API
###Usage  :
 
###Example:
  
###Flags  :
      --allow-self-signed-certs   

###Subcommands :
  

##set access

###Aliases:
  namespace-access, ns-access
###Usage  :
 set namespace access
###Example:
  chkit set access $USERNAME $ACCESS_LEVEL [--namespace $ID]
###Flags  :
  -f, --force   suppress confirmation

###Subcommands :
  

##run

###Aliases:
  
###Usage  :
 Run a solution
###Example:
  
###Flags  :
  -h, --help   Print help for chkit

###Subcommands :
  + solution : run solution from public template
  

##run solution

###Aliases:
  sol, solutions, sols, solu, so
###Usage  :
 run solution from public template
###Example:
  chkit run solution [$PUBLIC_SOLUTION] [--env=KEY1:VALUE1,KEY2:VALUE2] [--file $FILENAME] [--force]
###Flags  :
      --branch string   solution git repo branch, optional (default "master")
      --env string      solution environment variables, optional
      --file string     file with solution data, .yaml or .json, stdin if '-', optional
  -f, --force           create solution without confirmation, optional
      --name string     solution name, optional, autogenerated if void

###Subcommands :
  

##replace

###Aliases:
  
###Usage  :
 Replace deployment or service
###Example:
  
###Flags  :

###Subcommands :
  + configmap : 
  + deployment : replace deployment
  + ingress : patch ingress with new attributes
  + service : replace service
  

##replace service

###Aliases:
  srv, services, svc, serv
###Usage  :
 Replaces service.
Has an one-line mode, suitable for integration with other tools, and an interactive wizard mode
###Example:
  
###Flags  :
      --deployment string   deployment name, optional
      --domain string       service domain, optional
      --file string         create service from file
  -f, --force               suppress confirmation
      --port int            service external port, optional
      --port-name string    service port name
      --protocol string     service port protocol, optional (default "TCP")
      --target-port int     service target port, optional (default 80)

###Subcommands :
  

##replace ingress

###Aliases:
  ingr, ingresses, ing
###Usage  :
 Replaces ingress with new, use --force flag to write one-liner command, omitted attributes are inherited from previous ingress.
###Example:
  chkit replace ingress $INGRESS [--force] [--service $SERVICE] [--port 80] [--tls-secret letsencrypt]
###Flags  :
  -f, --force               replace ingress without confirmation
      --host string         ingress host, optional
      --port int            ingress endpoint port, optional (default 8080)
      --service string      ingress endpoin service, optional
      --tls-secret string   ingress tls-secret, use 'letsencrypt' for automatic HTTPS, '-' to use HTTP, optional

###Subcommands :
  

##replace deployment

###Aliases:
  depl, deployments, deploy
###Usage  :
 Replaces deployment.
Has an one-line mode, suitable for integration with other tools, and an interactive wizard mode
###Example:
  
###Flags  :
      --container-name string   container name, equal to image name by default
      --cpu uint                container CPU limit in mCPU, optional (default 200)
      --env stringArray         container env variable in KEY0:VALUE0 KEY1:VALUE1 format
      --file string             create deployment from file
  -f, --force                   suppress confirmation
      --image string            container image, optional
      --memory uint             container memory limit im Mb, optional (default 256)
      --replicas int            replicas, optional (default 1)

###Subcommands :
  

##replace configmap

###Aliases:
  cm, confmap, conf-map, comap
###Usage  :
 
###Example:
  
###Flags  :
      --file string         file with configmap data, .json, .yaml, .yml
      --file-item strings   configmap file item: $KEY:$FILENAME (default [])
      --force               suppress confirmation
      --item strings        configmap item: $KEY:$VALUE (default [])

###Subcommands :
  

##rename

###Aliases:
  
###Usage  :
 Rename resource
###Example:
  
###Flags  :

###Subcommands :
  + namespace : 
  

##rename namespace

###Aliases:
  ns, namespaces
###Usage  :
 
###Example:
  chkit rename ns $ID $NEW_NAME
###Flags  :

###Subcommands :
  

##logs

###Aliases:
  log
###Usage  :
 view pod logs. Aliases: log
###Example:
  logs pod_label [container] [--follow] [--prev] [--tail n] [--quiet]
###Flags  :
  -f, --follow      follow pod logs
  -t, --tail uint   print last <value> log lines (default 100)

###Subcommands :
  

##login

###Aliases:
  
###Usage  :
 Login to system
###Example:
  
###Flags  :
      --default-namespace string   use as default namespace, if '-', then use first one

###Subcommands :
  

##help [command]

###Aliases:
  
###Usage  :
 Help provides help for any command in the application.
Simply type chkit help [path to command] for full details.
###Example:
  
###Flags  :

###Subcommands :
  

##get

###Aliases:
  
###Usage  :
 Get resource data
###Example:
  
###Flags  :

###Subcommands :
  + access : get namespace access
  + configmap : 
  + containerum-api : 
  + default-namespace : print default
  + deployment : shows deployment data
  + ingress : show ingress data
  + namespace : shows namespace data or namespace list
  + pod : shows pod info
  + profile : show profile info
  + service : shows service info
  + solution : get solutions
  

##get solution

###Aliases:
  sol, solutions, sols, solu, so
###Usage  :
 Show solution list, available for run. To search solution by name add arg
###Example:
  chkit get solution [name]
###Flags  :

###Subcommands :
  

##get service

###Aliases:
  srv, services, svc, serv
###Usage  :
 chkit get service service_label [-o yaml/json] [-f output_file]
###Example:
  Shows service info
###Flags  :
  -f, --file string     output file (default "-")
  -o, --output string   output format [yaml/json]

###Subcommands :
  

##get profile

###Aliases:
  me, user
###Usage  :
 Shows profile info
###Example:
  chkit get profile
###Flags  :

###Subcommands :
  

##get pod

###Aliases:
  po, pods
###Usage  :
 shows pod info. Aliases: po, pods
###Example:
  chkit get pod pod_label [-o yaml/json] [-f output_file]
###Flags  :
  -f, --file string     output file (default "-")
  -o, --output string   output format (json/yaml)

###Subcommands :
  

##get namespace

###Aliases:
  ns, namespaces
###Usage  :
 shows namespace data or namespace list. Aliases: ns, namespaces
###Example:
  chkit get $ID... [-o yaml/json] [-f output_file]
###Flags  :
  -f, --file string     output file
  -o, --output string   output format (json/yaml)

###Subcommands :
  

##get ingress

###Aliases:
  ingr, ingresses, ing
###Usage  :
 Shows ingress data
###Example:
  chkit get ingress ingress_names... [-n namespace_label] [-o yaml/json]
###Flags  :
  -f, --file string     output file
  -o, --output string   output format (yaml/json)

###Subcommands :
  

##get deployment

###Aliases:
  depl, deployments, deploy
###Usage  :
 Shows deployment data
###Example:
  namespace deployment_names... [-n namespace_label]
###Flags  :
  -f, --file string     output file
  -o, --output string   output format (yaml/json)

###Subcommands :
  

##get default-namespace

###Aliases:
  default-ns, def-ns
###Usage  :
 print default
###Example:
  
###Flags  :

###Subcommands :
  

##get containerum-api

###Aliases:
  api, current-api, api-addr, API
###Usage  :
 
###Example:
  
###Flags  :

###Subcommands :
  

##get configmap

###Aliases:
  cm, confmap, conf-map, comap
###Usage  :
 
###Example:
  
###Flags  :
      --file string     output file (default "-")
  -o, --output string   output format yaml/json

###Subcommands :
  

##get access

###Aliases:
  namespace-access, ns-access
###Usage  :
 get namespace access
###Example:
  chkit get ns-access $ID
###Flags  :

###Subcommands :
  

##doc

###Aliases:
  
###Usage  :
 Print full chkit help
###Example:
  
###Flags  :
      --command string   print docs for command and its subcommands, example 'chkit doc --command "create depl"'
  -h, --help             help for doc
      --md               generate markdown docs
      --output string    output file, STDOUT by default

###Subcommands :
  

##delete

###Aliases:
  
###Usage  :
 Delete resource
###Example:
  
###Flags  :

###Subcommands :
  + configmap : delete configmap
  + deployment : delete deployment in specific namespace
  + ingress : delete ingress
  + namespace : delete namespace
  + pod : delete pod in specific namespace
  + service : delete service in specific namespace
  

##delete service

###Aliases:
  srv, services, svc, serv
###Usage  :
 Deletes service in namespace
###Example:
  chkit delete service service_label [-n namespace]
###Flags  :
  -f, --force   force delete without confirmation

###Subcommands :
  

##delete pod

###Aliases:
  po, pods
###Usage  :
 deletes pods. Aliases: po, pods
###Example:
  chkit delete pod pod_name [-n namespace]
###Flags  :
  -f, --force   delete pod without confirmation

###Subcommands :
  

##delete namespace

###Aliases:
  ns, namespaces
###Usage  :
 delete namespace deletes namespace with name, provided by first arg. Aliases: ns, namespaces
###Example:
  chkit delete namespace $ID
###Flags  :
  -f, --force   force delete without confirmation

###Subcommands :
  

##delete ingress

###Aliases:
  ingr, ingresses, ing
###Usage  :
 Deletes ingress
###Example:
  chkit delete ingress $INGRESS [-n $NAMESPACE] [--force]
###Flags  :
  -f, --force   delete ingress without confirmation

###Subcommands :
  

##delete deployment

###Aliases:
  depl, deployments, deploy
###Usage  :
 Deletes deployment in specific namespace.
Use --force flag to suppress confirmation
###Example:
  
###Flags  :
  -f, --force   delete without confirmation

###Subcommands :
  

##delete configmap

###Aliases:
  cm, confmap, conf-map, comap
###Usage  :
 delete configmap
###Example:
  
###Flags  :
  -f, --force   suppress confirmation

###Subcommands :
  

##create

###Aliases:
  
###Usage  :
 Create deployment or service
###Example:
  
###Flags  :

###Subcommands :
  + configmap : 
  + deployment : create deployment
  + ingress : create ingress
  + service : create service
  

##create service

###Aliases:
  srv, services, svc, serv
###Usage  :
 create service for provided pod in provided namespace
###Example:
  
###Flags  :
      --deploy string      service deployment, required
  -f, --file string        file with service data (default "-")
      --force              create service without confirmation
      --name string        service name, optional (default "sage-geiger")
      --port int           service port, optional
      --port-name string   service port name, optional (default "amundsenia-manatee")
      --proto string       service protocol, optional (default "TCP")
      --target-port int    service target port, optional (default 80)

###Subcommands :
  

##create ingress

###Aliases:
  ingr, ingresses, ing
###Usage  :
 Creates ingress. TLS with LetsEncrypt and custom cert is available
###Example:
  chkit create ingress [--force] [--filename ingress.json] [-n prettyNamespace]
###Flags  :
  -f, --force               create ingress without confirmation
      --host string         ingress host (example: prettyblog.io), required
      --path string         path to endpoint (example: /content/pages), optional
      --port int            ingress endpoint port (example: 80, 443), optional (default 8080)
      --service string      ingress endpoint service, required
      --tls-cert string     TLS cert file, optional
      --tls-secret string   TLS secret string, optional

###Subcommands :
  

##create deployment

###Aliases:
  depl, deployments, deploy
###Usage  :
 Creates new deployment.
Has an one-line mode, suitable for integration with other tools,
and an interactive wizard mod
###Example:
  
###Flags  :
      --configmap strings   container configmap, CONTAINER_NAME@CONFIGMAP_NAME@MOUNTPATH in case of multiple containers or
                            CONFIGMAP_NAME@MOUNTPATH or CONFIGMAP_NAME in case of one container.
                            If MOUNTPATH is omitted, then use /etc/CONFIGMAP_NAME as mountpath (default [])
      --cpu strings         container memory limit, mCPU,
                            CONTAINER_NAME@CPU in case of multiple containers or CPU in case of one container (default [])
      --env strings         container environment variable,
                            CONTAINER_NAME@KEY:VALUE in case of multiple containers or KEY:VALUE in case of one container (default [])
      --file string         file with configmap data, .json, .yaml, .yml, optional
  -f, --force               suppress confirmation, optional
      --image strings       container image,
                            CONTAINER_NAME@IMAGE in case of multiple containers or IMAGE in case of one container (default [])
      --memory strings      container memory limit, Mb,
                            CONTAINER_NAME@MEMORY in case of multiple containers or MEMORY in case of one container (default [])
      --name string         deployment name, optional
      --replicas uint       deployment replicas, optional
      --volume strings      container volume,
                            CONTAINER_NAME@VOLUME_NAME@MOUNTPATH in case of multiple containers or
                            VOLUME_NAME@MOUNTPATH or VOLUME_NAME in case of one container.
                            If MOUNTPATH is omitted, then use /mnt/VOLUME_NAME as mountpath (default [])

###Subcommands :
  

##create configmap

###Aliases:
  cm, confmap, conf-map, comap
###Usage  :
 
###Example:
  
###Flags  :
      --file string           file with configmap data
  -f, --force                 suppress confirmation
      --item-file strings     configmap file, KEY:FILE_PATH or FILE_PATH
      --item-string strings   configmap item, KEY:VALUE string pair
      --name string           configmap name (default "phocaea-jacobi")

###Subcommands :
  
