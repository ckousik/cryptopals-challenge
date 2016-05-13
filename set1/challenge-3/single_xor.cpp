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
        bin += (char)(16 * hex_val(hex[i]) + hex_val(hex[i+1]));
    }
    return bin;
}

int score_char(char x){
    string letter_scores = "etaonrishd .,\nlfcmugypwbvkjxqz-_!?'\"/1234567890*";
    int score = letter_scores.find(x);
    if(score == -1){
        return score;
    }
    return 256-score;
}

string single_xor(string s, char& key){
    int probable_candidates[256];
    string bin = hex_decode(s);
    int i;
    for(i = 0; i < 256; i++){
        int count = 0;
        for(int j = 0; j< bin.length(); j++){
            if(bin[j] == (char)i) count++;
        }
        probable_candidates[i] = count * score_char(i);
    }
    int m = probable_candidates[0];
    char mi = 0;
    for(i = 0; i<256;i++){
        if(probable_candidates[i] > m){
            m = probable_candidates[i];
            mi = i;
        }
    }
    string out;
    for(int i = 0;i<bin.length();i++){
        out += bin[i]^mi;
    }
    key = mi;
    return out;
}

string xorKey (string s, char key){
    string out;
    for(int i=0;i<s.length();i++){
        out += key^s[i];
    }
    return out;
}

int main(int argc, char* argv[])
{
    char k;
    string input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";
    string output = single_xor(input,k);
    cout<<"Key:"<<k<<'\n'<<output.c_str()<<'\n';
    return 0;
}
