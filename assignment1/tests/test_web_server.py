import unittest
import requests

class TestWebServer(unittest.TestCase):
    def test_http_redirects_to_https(self):
        response = requests.get('http://www.mywebsite.com', allow_redirects=False)
        self.assertEqual(response.status_code, 302)
        self.assertEqual(response.headers['Location'], 'https://www.mywebsite.com/')

    def test_https_content(self):
        response = requests.get('https://www.mywebsite.com')
        self.assertEqual(response.status_code, 200)
        self.assertIn('Welcome to my answer for the sre challange given as part of the interview process', response.text)

    def test_https_error_handling(self):
        response = requests.get('https://www.mywebsite.com/non-existent-url')
        self.assertEqual(response.status_code, 404)

if __name__ == '__main__':
    unittest.main()
