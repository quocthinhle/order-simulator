export type HttpRequestAcceptableDataType = string | number | string[] | number[];

export type Queryable = Record<string, HttpRequestAcceptableDataType>;

export type IHttpRequest = {
    url: string;
    query?: Queryable,
    body?: Record<string, HttpRequestAcceptableDataType>,
    headers?: Record<string, string>, // Cause header only accept string data-type.
}