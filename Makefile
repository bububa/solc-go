deps: libboost libsolc

libboost: 
	tar -vxf ./dependencies/boost_1_67_0.tar.gz -C ./dependencies;
	cd ./dependencies/boost_1_67_0;sh ./bootstrap.sh; 
	./b2 --with-regex --with-filesystem variant=release link=static stage; cd ../../;
	cp ./dependencies/boost_1_67_0/stage/lib/libboost_filesystem.a ./lib/;
	cp ./dependencies/boost_1_67_0/stage/lib/libboost_regex.a ./lib/;

libsolc:
	git clone -b release https://github.com/ethereum/solidity.git ./dependencies/solidity;
	sh ./dependencies/solidity/scripts/install_deps.sh;
	sh ./dependencies/solidity/scripts/build.sh;
	cp ./dependencies/solidity/build/libsolc/libsolc.a ./lib/;
	cp ./dependencies/solidity/libsolc/libsolc.h ./include/;
	cp ./dependencies/solidity/build/libdevcore/libdevcore.a ./lib/;
	cp ./dependencies/solidity/build/libevmasm/libevmasm.a ./lib/;
	cp ./dependencies/solidity/build/libsolidity/libsolidity.a ./lib/;
	cp -r ./dependencies/solidity/build/deps/include/* ./include;
	cp -r ./dependencies/solidity/build/deps/lib/* ./lib/;