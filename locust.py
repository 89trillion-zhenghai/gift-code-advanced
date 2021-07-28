from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def createAndGetGiftCode01(self):
       self.client.post('/createAndGetGiftCode',data={'userName':'admin01','description':'十周年活动奖励','giftType':'2','validity':'2m','availableTimes':'9999','giftDetail':'{"1001":"10","1002":"20"}'})

    @task
    def login01(self):
       self.client.post('/login',data={'name':'testName01'})

    @task
    def login02(self):
       self.client.post('/login',data={'name':'testName02'})

    @task
    def login03(self):
      self.client.post('/login',data={'name':'testName03'})

    @task
    def login04(self):
       self.client.post('/login',data={'name':'testName04'})

    @task
    def redeemGift01(self):
       self.client.post('/redeemGift',data={'name':'testName01','giftCode':'8K187418'})

    @task
    def redeemGift02(self):
       self.client.post('/redeemGift',data={'name':'testName02','giftCode':'8K187418'})

    @task
    def redeemGift03(self):
       self.client.post('/redeemGift',data={'name':'testName03','giftCode':'8K187418'})

    @task
    def getGiftDetail01(self):
       self.client.post('/getGiftDetail',data={'giftCode':'JF362262'})
