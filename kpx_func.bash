# kpx_func.sh -- shell-function for BASH/ZSH for k8s-commands for Portworx
#

function kpx() {
	local sw="$1" 
	local cmd
	local usage='
Usage: kpx <command>

command:
   ls               List all Portworx pods
   w[atch]          Watch the status of Portworx pods
   des[cribe]       Describe Portworx DaemonSet
   ed[it]           Edit Portworx DaemonSet
   del[ete]         Remove Portworx deployment from Kubernetes (note: Portworx service unaffected)
   rol[lout]        Monitor Portworx update rollout
   log <px-pod>     Displays a log of a Portworx pod
   lf <px-pod>      Follows log of a Portworx pod
   kill <px-pods>   Deletes Portworx pod (force + no delay)
   cordon <node>    Cordons nodes (use --all for all nodes)
   uncordon <node>  UnCordons nodes (use --all for all nodes)
   lab[el] [<false|remove> [<node1> [<node2>...]]]
                    Labels the nodes with given label (no params - removes label from all nodes)
   config           Dumps current kubectl config (store into $HOME/.kube/config)
' 
	[ $# -lt 1 ] && echo "$usage" && return 0
	shift
	case "$sw" in
		(ls) cmd="kubectl get pods -o wide -n kube-system -l name=portworx $@"  ;;
		(w*) cmd="watch kubectl get pods -o wide -n kube-system -l name=portworx $@"  ;;
		(des*) cmd="kubectl describe -n kube-system ${@:-ds/portworx}"  ;;
		(ed*) cmd="kubectl edit -n kube-system ${@:-ds/portworx}"  ;;
		(del*) cmd="kubectl delete -f 'http://install.portworx.com?stork=true&ctl=true&lh=true$@'"  ;;
		(rol*) cmd="kubectl rollout status -n kube-system ${@:-ds/portworx}"  ;;
		(lab*) cmd=- 
			[ x"$1" != x ] && cmd="=$1"  && shift
			cmd="kubectl label nodes ${@:---all} px/enabled$cmd"  ;;
		(log*) cmd="kubectl -n kube-system logs $@"  ;;
		(lf) cmd="kubectl -n kube-system logs -f $@"  ;;
		(kill) cmd="kubectl -n kube-system delete pod --grace-period=0 --force $@"  ;;
		(conf*) cmd="kubectl config view --flatten --minify $@"  ;;
		(exe*) cmd="${1:-`kubectl get pods -n kube-system -l name=portworx -o name | cut -d/ -f2 | head -1`}" 
			shift 2> /dev/null
			cmd="kubectl -n kube-system exec -it $cmd ${@:-/bin/bash}"  ;;
		(cordon|uncordon) if [ x"$1" = x--all ]
			then
				cmd="kubectl get nodes -o name | xargs -n1 kubectl $sw" 
			else
				cmd="kubectl $sw $@" 
			fi ;;
		(-h*) echo "$usage"
			return 1 ;;
		(*) echo "$0: Unknown option '$1'" >&2
			return 1 ;;
	esac
	echo -e "\033[7m >> $cmd \033[m" >&2
	eval "$cmd"
}

