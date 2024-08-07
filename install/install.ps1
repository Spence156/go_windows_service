[string]$serviceName = 'example_service'
[string]$installerPath = 'C:\repos\personal\go_windows_service\example_service.exe'

sc.exe delete $serviceName
sc.exe create $serviceName binpath=$installerPath
sc.exe description $serviceName "This is an example service built using golang."
