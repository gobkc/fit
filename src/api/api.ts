import type {
    Response,
    ApiVersionResponse,
    ListCateResponse,
    ListNoteResponse,
    NewNoteRequest,
    ListConfigurationsResponse, AppConfig
} from "./interface"
import {http} from "./config";

export class Api {
    private version: number = 0;

    apiVersion(): Promise<ApiVersionResponse> {
        return http({
            url: `/p/version`,
            method: 'get',
            data: {}
        }).then(response => response.data as ApiVersionResponse);
    }

    listCate(): Promise<ListCateResponse> {
        return http({
            url: `/p/list-cate`,
            method: 'get',
            data: {}
        }).then(response => response.data as ListCateResponse);
    }

    ListNote(keyword: string): Promise<ListNoteResponse> {
        return http({
            url: `/p/list-note?keyword=${keyword}`,
            method: 'get',
            data: {}
        }).then(response => response.data as ListNoteResponse);
    }

    ListCateNote(cate: string, keyword: string): Promise<ListNoteResponse> {
        return http({
            url: `/p/${cate}/list-note?keyword=${keyword}`,
            method: 'get',
            data: {}
        }).then(response => response.data as ListNoteResponse);
    }

    NewCate(cate: string): Promise<Response> {
        return http({
            url: `/p/new-cate`,
            method: 'post',
            data: {cate: cate}
        }).then(response => response.data as Response);
    }

    NewNote(request: NewNoteRequest): Promise<Response> {
        return http({
            url: `/p/new-note`,
            method: 'post',
            data: request
        }).then(response => response.data as Response);
    }

    Pull(): Promise<Response> {
        return http({
            url: `/p/pull`,
            method: 'post',
            data: {}
        }).then(response => response.data as Response);
    }

    Push(): Promise<Response> {
        return http({
            url: `/p/push`,
            method: 'post',
            data: {}
        }).then(response => response.data as Response);
    }

    listConfigurations(): Promise<AppConfig[]> {
        return http({
            url: `/p/list-conf`,
            method: 'get',
            data: {}
        }).then(response => response.data as AppConfig[]);
    }
}
