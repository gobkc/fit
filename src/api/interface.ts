export interface Response {
    error: number,
    msg: string,
    more: string,
}

export interface ApiVersionResponse {
    apiVersion: string
}

export interface ListCateResponse {
    error: number,
    msg: string,
    more: string,
    data: string[],
}

export interface Note {
    content: number,
    title: string,
    updated_time: string,
}

export interface ListNoteResponse {
    error: number,
    msg: string,
    more: string,
    data: Note[],
}


export interface NewNoteRequest {
    cate: string,
    content: string,
    title: string,
}

export interface EmailConfig {
    imap: string;
    smtp: string;
    user: string;
    pass: string;
}

export interface CorsConfig {
    enabled: boolean;
    max_age: number;
    allowed_origins: string[];
    allowed_methods: string[];
    allowed_headers: string[];
    allow_credentials: boolean;
}

export interface AppConfig {
    name: string;
    version: string;
    debug: boolean;
    rest_addr: string;
    dsn: string;
    email: EmailConfig;
    cors: CorsConfig;
    max_idle: number;
    max_conn: number;
    max_left_time: number;
    jwt_salt: string;
}

export interface ListConfigurationsResponse {
    error: number,
    msg: string,
    more: string,
    data: AppConfig[],
    main_conf: string,
}
