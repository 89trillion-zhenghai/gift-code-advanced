from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def admin_post01(self):
       self.client.get('/login?name=testName')

    @task
    def admin_post02(self):
       self.client.get('/redeemGift?name=test&giftCode=33N3110J',)

    @task
    def admin_post03(self):
       self.client.get('/redeemGift?name=testName&giftCode=33N3110J',)