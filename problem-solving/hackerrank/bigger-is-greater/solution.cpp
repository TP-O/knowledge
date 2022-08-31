#include <bits/stdc++.h>

using namespace std;

// Complete the biggerIsGreater function below.
string biggerIsGreater(string w) {
    for (int i = w.length() - 1; i > 0; i--) {
        if (w[i - 1] < w[i]) {
            for (int j = w.length() - 1; j >= i; j--) {
                if (w[i - 1] < w[j]) {
                    swap(w[i - 1], w[j]);

                    break;
                }
            }

            string l = w.substr(0, i);
            string r = w.substr(i, w.length());
            reverse(r.begin(), r.end());

            return l + r;
        }
    }

    return "no answer";
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int T;
    cin >> T;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int T_itr = 0; T_itr < T; T_itr++) {
        string w;
        getline(cin, w);

        string result = biggerIsGreater(w);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}
