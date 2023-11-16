import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpRequest  } from '@angular/common/http';
import { IHttpRequest, Queryable } from './types/request';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class HttpCommonService {
    constructor(
        private builtInHttpService: HttpClient,
    ) { }

    private serializeQueryParams(queries: Queryable): string {
        return (new HttpParams(queries)).toString();
    }

    get(request: IHttpRequest): Observable<any> {
        if (request.query) {
            request.url = `${request.url}?${this.serializeQueryParams(request.query)}`;
        }

        return this.builtInHttpService
            .get(request.url);
    }

    post(request: IHttpRequest) {
        if (request.query) {
            request.url = `${request.url}?${this.serializeQueryParams(request.query)}`;
        }

        return this.builtInHttpService
            .post(
                request.url,
                request.body,
            );
    }

    patch() {

    }

    put() {

    }

    delete() {}
}