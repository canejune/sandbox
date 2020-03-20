#include <iostream>
#include <fstream>      // std::ifstream
#include <algorithm>
#include <vector>
#include <string>

#include <chrono> 
using namespace std::chrono; 
using namespace std;

void use_getline(ifstream &is, vector <string> &vLines)
{
    string line;
    while (getline(is, line))
    {
        vLines.push_back(line);
    }
}

void use_bufferRead(ifstream &is, vector <long long> &vLineOffsets)
{
    string line;
    static const int bufSize = 1024 * 1024;
    char *buffer = new char [bufSize];
    long long idx, nRead, offset = 0;
    
    vLineOffsets.push_back(0);
    while(is.good()) {
        is.read(buffer, bufSize);
        nRead = is.gcount();
         for (idx = 0; idx < nRead; idx++) {
            if (buffer[idx] == '\n') {
                vLineOffsets.push_back(offset + idx + 1);
            }
        }
        offset += idx;
    }
    delete[] buffer;
}


int main(void)
{
    static const char* bigFilePath = "C:\\Users\\canej\\Study\\testfiles\\100M.log";
    
    ifstream bigFile;
    vector <string> vLines;
    vector <long long> vLineOffsets;

    bigFile.open(bigFilePath);
    auto start = high_resolution_clock::now();
    use_getline(bigFile, vLines);
    auto stop = high_resolution_clock::now();
    auto duration = duration_cast<milliseconds>(stop - start); 
    cout << duration.count() << " " << vLines.size() << endl;
    bigFile.close();

    cout << vLines[3] << endl;
    cout << vLines[3].length() << endl;
    vLines.clear();

    bigFile.open(bigFilePath, ifstream::binary);
    start = high_resolution_clock::now();
    use_bufferRead(bigFile, vLineOffsets);
    stop = high_resolution_clock::now();
    duration = duration_cast<milliseconds>(stop - start); 
    cout << duration.count() << " " << vLineOffsets.size() << endl;

    bigFile.clear();

    bigFile.seekg(vLineOffsets[3]);

    char buffer[256] = {0,};
    bigFile.read(buffer, vLineOffsets[4] - vLineOffsets[3]);

    cout << buffer << endl;

    bigFile.close();

    return 0;
}