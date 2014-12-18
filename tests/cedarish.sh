
readonly cedarish_version="v2"
readonly cedarish_image="progrium/cedarish"
readonly cedarish_url="https://github.com/progrium/cedarish/releases/download/$cedarish_version/cedarish-cedar14_$cedarish_version.tar.gz"

check-cedarish() {
	docker images | grep "$cedarish_image" | grep "$cedarish_version" > /dev/null
}

import-cedarish() {
	local version imagetag
	version="$(docker version | head -1 | cut -d' ' -f 3)"
	# CircleCI is running a fork of Docker 1.2
	if [[ "${version:0:3}" == "1.2" ]]; then
		imagetag="$cedarish_image $cedarish_version"
	else
		imagetag="$cedarish_image:$cedarish_version"
	fi
	if [[ -f ".cache/cedarish_$cedarish_version.tgz" ]]; then
		time cat ".cache/cedarish_$cedarish_version.tgz" | docker import - $imagetag
	else
		time docker import "$cedarish_url" $imagetag
	fi
	echo
}