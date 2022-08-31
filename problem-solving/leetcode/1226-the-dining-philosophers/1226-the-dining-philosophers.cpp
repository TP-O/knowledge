using namespace std;

class DiningPhilosophers {

private:
    int pickedForked[5] = {0, 0, 0, 0, 0};
    condition_variable cv;
    mutex mtx;
    
public:
    DiningPhilosophers() {        
        //
    }

    void wantsToEat(int philosopher,
                    function<void()> pickLeftFork,
                    function<void()> pickRightFork,
                    function<void()> eat,
                    function<void()> putLeftFork,
                    function<void()> putRightFork) {
        unique_lock<mutex> lk(mtx);
        
        cv.wait(lk, [&]() { 
           return pickedForked[(philosopher) % 5] == 0; 
        });
        
        pickLeftFork();
        
        cv.wait(lk, [&]() { 
           return pickedForked[(philosopher + 1) % 5] == 0; 
        });
        
        pickRightFork();
        
        eat();
        
        putLeftFork();
        putRightFork();
    }
};