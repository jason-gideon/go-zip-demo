#include <cstdio>
#include "uniq_ptr.h"
#include "sharedptrr.h"
#include "move_construct.h"

using namespace std;


#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include <zlib.h>


int libz()
{
	char inbuf[] = "Hello, This is a demo for compress and uncompress interface!\n"
		"Written by windeal.li\n"
		"email: 2378264731@qq.com\n";
	uLong inlen = sizeof(inbuf);
	char *outbuf = NULL;
	uLong outlen;

	outlen = compressBound(inlen);
	printf("in_len: %ld\n", inlen);
	printf("out_len: %ld\n", outlen);

	if ((outbuf = (char *)malloc(sizeof(char) * outlen)) == NULL) {
		fprintf(stderr, "Failed to malloc memory for outbuf!\n");
		return -1;
	}

	/* compress */
	if (compress(outbuf, &outlen, inbuf, inlen) != Z_OK) {
		fprintf(stderr, "Compress failed!\n");
		return -1;
	}
	printf("Compress Sucess!\n");
	printf("\routlen:%ld, outbuf:%s\n", outlen, outbuf);

	memset(inbuf, 0, sizeof(inbuf));
	/* Uncompress */
	if (uncompress(inbuf, &inlen, outbuf, outlen) != Z_OK) {
		fprintf(stderr, "Uncompress failed!\n");
		return -1;
	}
	printf("Uncompress Success!\n");
	printf("\rinlen:%ld, inbuf:%s\n", inlen, inbuf);

	/* free memory */
	if (outbuf != NULL) {
		free(outbuf);
		outbuf = NULL;
	}

	return 0;
}




string a("sss");
string& getStr() {
	
	return a;
}


int intaa = 10;
const int& getInt() {
	return intaa;
}


int main()
{
    printf("hello from cxx11!\n");

	libz();
	return 0;  //zip test


	string& a = getStr();
	string b = getStr();


	const int & c = getInt();
	//c = 100;

	printf("%d\n", intaa);


	//1. uniq_ptr
	//独享对象，不能赋值，只能move
	/*reset() 重置unique_ptr为空，delete其关联的指针。
	  release() 不delete关联指针，并返回关联指针。释放关联指针的所有权，unique_ptr为空。
	  get() 仅仅返回关联指针 */
	{
		//方法1： 构建uniquptr
		unique_ptr<task> t1(new task(10));
		//方法2：构建
		//unique_ptr<task> t2 = make_unique<task>(20);    //cxx14下才有make_unique
		//创建错误
		//unique_ptr<task> t3 = new task();    //cxx14下才有make_unique

		//通过 unique_ptr 访问其成员
		int id = t1->id;
		std::cout << id << std::endl;

		////////扩展1//
		//创建一个nil的unique对象
		unique_ptr<int> ptr1;

		////////扩展2//
		//检查unique对象是不是空
		// 方法1
		if (!ptr1)
			std::cout << "ptr1 is empty" << std::endl;
		// 方法2
		if (ptr1 == nullptr)
			std::cout << "ptr1 is empty" << std::endl;

		////////扩展3//
		//获取被管理的指针
		task *p1 = t1.get();



		////////扩展5//
		//转移unique_ptr对象所有权
		unique_ptr<task> t3 = move(t1);
		if (t1 == nullptr){
			std::cout << "t1 is  empty" << std::endl;
		}


		////////扩展4//
		//重置uniqptr。reset会触发delete释放维护的对象
		t1.reset();
	}
	

	//2. shared_ptr
	{
		//test_sharedptr();
	}

	//3.移动构造函数
	{
		test_move_cons();
	}

    return 0;
}