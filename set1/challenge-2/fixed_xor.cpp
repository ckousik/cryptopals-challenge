#include <iostream>
#include <string>
#include <cassert>

using namespace std;

const char hex_table[] = "0123456789abcdef";

int hex_val(char c){
    tolower(c);
    for(int i = 0; i<16; i++){
        if(hex_table[i] == c) return i;
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

string hex_encode(string bin){
    string hex;
    int size = bin.length(), i;
    for(i=0; i< size;i++){
        char h = bin[i];
        hex += hex_table[(h & 0xF0)>>4];
        hex += hex_table[(h & 0x0F)];
    }
    return hex;
}

string fixed_xor(string a, string b){
    assert(a.length() == b.length());
    a = hex_decode(a);
    b = hex_decode(b);
    int size = a.length();
    string c;
    for(int i=0;i<size;i++){
        c += (char) a[i]^b[i];
    }
    return hex_encode(c);
}

int main(int argc, char* argv[])
{
    string input1 = "1c0111001f010100061a024b53535009181c",
    input2 = "686974207468652062756c6c277320657965",
    expected_output = "746865206b696420646f6e277420706c6179",
    output = fixed_xor(input1, input2);
    cout<<output.c_str()<<'\n';
    assert(output.compare(expected_output) == 0);
    return 0;
}
