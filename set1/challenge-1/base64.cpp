#include <iostream>
#include <string>
#include <cassert>

using namespace std;

int hex_val(char c){
    const char hex[] = "0123456789abcdef";
    tolower(c);
    for(int i = 0; i<16; i++){
        if(hex[i] == c) return i;
    }
    return -1;
}


string hex_decode (string hex){
    string bin;
    int size16 = hex.length(), i;
    for(i = 0;i< size16;i+=2){
        bin += (unsigned char)(16 * hex_val(hex[i]) + hex_val(hex[i+1]));
    }
    return bin;
}

string hex_to_b64 (string hex){
    const string enc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    string b64 , bin = hex_decode(hex);
    int i;

    // const char[] extract = {
    //     0xFC,    11111100 a
    //     3,       00000011 a
    //     0xF0,    11110000 b
    //     0x0F,    00001111 b
    //     0xC0,    11000000 c
    //     0x3F     00111111 c
    // };

    int pad = bin.length()%3;
    for(i = pad; i > 0; i++){
        bin+= (char)0x00;
    }
    for(i = 0; i<bin.length(); i+=3){
        int v[4];
        char a,b,c;
        a = bin[i]; b = bin[i+1]; c = bin[i+2];
        v[0] = (a & 0xFC) >> 2;
        v[1] = ((a & 3) << 4) | ((b & 0xF0) >> 4);
        v[2] = ((b & 0x0F) << 2) | ((c & 0xC0) >> 6);
        v[3] = (c & 0x3F);
        for(int j = 0; j<4; j++){
            b64 += enc[v[j]];
        }
    }
    for(i = 0; i < pad; i++) b64+='=';
    return b64;
}

int main(int argc, char* argv[])
{
    string input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
    expected_output = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
    output = hex_to_b64(input);
    cout<<output.c_str()<<'\n';
    assert(output.compare(expected_output) == 0);
    return 0;
}
