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
	//������󣬲��ܸ�ֵ��ֻ��move
	/*reset() ����unique_ptrΪ�գ�delete�������ָ�롣
	  release() ��delete����ָ�룬�����ع���ָ�롣�ͷŹ���ָ�������Ȩ��unique_ptrΪ�ա�
	  get() �������ع���ָ�� */
	{
		//����1�� ����uniquptr
		unique_ptr<task> t1(new task(10));
		//����2������
		//unique_ptr<task> t2 = make_unique<task>(20);    //cxx14�²���make_unique
		//��������
		//unique_ptr<task> t3 = new task();    //cxx14�²���make_unique

		//ͨ�� unique_ptr �������Ա
		int id = t1->id;
		std::cout << id << std::endl;

		////////��չ1//
		//����һ��nil��unique����
		unique_ptr<int> ptr1;

		////////��չ2//
		//���unique�����ǲ��ǿ�
		// ����1
		if (!ptr1)
			std::cout << "ptr1 is empty" << std::endl;
		// ����2
		if (ptr1 == nullptr)
			std::cout << "ptr1 is empty" << std::endl;

		////////��չ3//
		//��ȡ�������ָ��
		task *p1 = t1.get();



		////////��չ5//
		//ת��unique_ptr��������Ȩ
		unique_ptr<task> t3 = move(t1);
		if (t1 == nullptr){
			std::cout << "t1 is  empty" << std::endl;
		}


		////////��չ4//
		//����uniqptr��reset�ᴥ��delete�ͷ�ά���Ķ���
		t1.reset();
	}
	

	//2. shared_ptr
	{
		//test_sharedptr();
	}

	//3.�ƶ����캯��
	{
		test_move_cons();
	}

    return 0;
}