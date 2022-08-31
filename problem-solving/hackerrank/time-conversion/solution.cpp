#include <bits/stdc++.h>
#include <string>

using namespace std;

string get_tail(string &s) {
    string tail = "";

    tail = s.substr(s.size() - 2);
    s.erase(s.size() - 2, s.size() - 1);

    return tail;
}

string get_head(string &s) {
    string head = "";

    head = s.substr(0, 2);
    s.erase(0, 2);

    return head;
}
/*
 * Complete the timeConversion function below.
 */
string timeConversion(string s) {
    /*
     * Write your code here.
     */
    string tail = get_tail(s);
    string head = get_head(s);
    int hour = atoi(head.c_str());

    if(tail == "AM") {
        if(hour == 12) {
            head = "00";
        }
    } else if(tail == "PM") {
        if(hour != 12) {
            hour += 12;
            head = to_string(hour);
        }
    } else {
        return "ERROR!";
    }

    s.insert(0, head);

    return s;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    string result = timeConversion(s);

    fout << result << "\n";

    fout.close();

    return 0;
}
