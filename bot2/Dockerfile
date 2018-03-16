FROM ubuntu:14.04.5
ADD victim.sh /
ADD init.sh /
RUN chmod +x init.sh \
	&& chmod +x victim.sh
RUN apt-get -y update && /victim.sh

	
EXPOSE 23
