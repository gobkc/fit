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

    listNote(keyword: string): Promise<ListNoteResponse> {
        return http({
            url: `/p/list-note?keyword=${keyword}`,
            method: 'get',
            data: {}
        }).then(response => response.data as ListNoteResponse);
    }

    listCateNote(cate: string, keyword: string): Promise<ListNoteResponse> {
        return http({
            url: `/p/${cate}/list-note?keyword=${keyword}`,
            method: 'get',
            data: {}
        }).then(response => response.data as ListNoteResponse);
    }

    newCate(cate: string): Promise<Response> {
        return http({
            url: `/p/new-cate`,
            method: 'post',
            data: {cate: cate}
        }).then(response => response.data as Response);
    }

    deleteCate(cate: string): Promise<Response> {
        return http({
            url: `/p/cate`,
            method: 'delete',
            data: {cate: cate}
        }).then(response => response.data as Response);
    }

    deleteNote(cate: string, title: string): Promise<Response> {
        return http({
            url: `/p/note`,
            method: 'delete',
            data: {cate: cate, title: title}
        }).then(response => response.data as Response);
    }

    newNote(request: NewNoteRequest): Promise<Response> {
        return http({
            url: `/p/new-note`,
            method: 'post',
            data: request
        }).then(response => response.data as Response);
    }

    pull(): Promise<Response> {
        return http({
            url: `/p/pull`,
            method: 'post',
            data: {}
        }).then(response => response.data as Response);
    }

    push(): Promise<Response> {
        return http({
            url: `/p/push`,
            method: 'post',
            data: {}
        }).then(response => response.data as Response);
    }

    listConfigurations(): Promise<ListConfigurationsResponse> {
        return http({
            url: `/p/list-conf`,
            method: 'get',
            data: {}
        }).then(response => response.data as ListConfigurationsResponse);
    }

    deleteConfigurations(conf_name: string): Promise<Response> {
        return http({
            url: `/p/conf`,
            method: 'delete',
            data: {
                name: conf_name
            }
        }).then(response => response.data as Response);
    }

    createConfiguration(conf: AppConfig): Promise<Response> {
        return http({
            url: `/p/create-conf`,
            method: 'post',
            data: {"conf": conf,}
        }).then(response => response.data as Response);
    }

    enableConfiguration(conf: AppConfig): Promise<Response> {
        return http({
            url: `/p/enable-conf`,
            method: 'post',
            data: {"conf": conf,}
        }).then(response => response.data as Response);
    }
}
