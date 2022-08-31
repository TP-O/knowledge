#include <bits/stdc++.h>

using namespace std;

// Complete the staircase function below.
void staircase(int n, int p) {
    if(n == p) {
        for(int i = 0; i < n; i++)
            cout << '#';
        return;
    }

    for(int i = 1; i <= n; i++) {
        if(i > n-p)
            cout << '#';
        else
            cout << ' ';
    }
    cout << '\n';

    return staircase(n, p+1);
}

int main()
{
    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    staircase(n, 1);

    return 0;
}
