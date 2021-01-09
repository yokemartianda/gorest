import http from 'k6/http';
import { group,check, sleep } from 'k6';
export let options = {
    vus: 10,
    duration: '30s',
};

const SLEEP_DURATION = 0.1;
const BASE_URL = "http://localhost:8080/api/v1/";

export default function () {
    let body = JSON.stringify({
        name: 'user_' + __ITER,
        address: 'address_' + __ITER,
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    group('simple customer journey', (_) => {
        // Create customer request
        let login_response = http.post(
            BASE_URL + "customers",
            body,
            params,
        );
        check(login_response, {
            'is status 201': (r) => r.status === 201,
            'is message present': (r) => r.json().hasOwnProperty('message'),
        });
        sleep(SLEEP_DURATION);

        let customer_response = http.get(
            BASE_URL + "customers/" + __ITER,
            params,
        );
        check(customer_response, {
            'is status 200': (r) => r.status === 200,
        });
        sleep(SLEEP_DURATION);

        // Update customer request
        body = JSON.stringify({
            name: 'user_' + __ITER,
            address: 'address_' + __ITER,
        });
        let update_customer_response = http.post(
            BASE_URL + "customers/" + __ITER,
            body,
            params,
        );
        check(update_customer_response, {
            'is status 202': (r) => r.status === 202,
        });
        sleep(SLEEP_DURATION);
    });
}