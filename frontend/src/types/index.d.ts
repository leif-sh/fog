// 基础
export interface Meta {
  views: number;
  likes: number;
  comments: number;
}

export interface ToUser {
  user_id: number;
  name: string;
  avatar: string;
  type: number;
}
export interface Params {
  keyword: string;
  pageNum: number;
  pageSize: number;
}

// 登录
export interface LoginParams {
  email: string;
  password: string;
}
export interface UserInfo {
  id: number;
  name: string;
  avatar: string | any;
}
export interface RegAndLogParams {
  email: string;
  name: string;
  password: string;
  phone: string;
  desc: string;
}


// 导航 nav
export interface NavListItem {
  index: string;
  path: string;
  name: string;
}



// 文章归档
export interface ParamsArchive {
  keyword: string;
  likes: string; // 是否是热门文章
  state: number; // 文章发布状态 => 0 草稿，1 已发布,'' 代表所有文章
  article: number;
  tag_id: string;
  category_id: string;
  pageNum: number;
  pageSize: number;
}
export interface ArchiveListItem {
  created_at: string;
  title: string;
  id: string;
}
export interface ArchiveList {
  year: string;
  list: ArchiveListItem[];
}
export interface ArchiveData {
  count: number;
  list: ArchiveList | any;
}


// 文章详情
export interface OtherComments {
  content: string;
  created_at: string;
  likes: number;
  state: number;
  to_user: ToUser;
  user: ToUser;
  id: number;
}
export interface Comments {
  article_id: number;
  content: string;
  created_at: string;
  id: number;
  is_handle: number;
  is_top: boolean;
  likes: number;
  other_comments: OtherComments[];
  state: number;
  updated_at: number;
  user: ToUser;
  user_id: string;
  __v: number;
}

export interface ArticleDetailIF {
  toc: string;
  author: string;
  category: Array<any>;
  comments: Array<Comments>;
  created_at: number;
  desc: string;
  content: string;
  id: number;
  img_url: string;
  numbers: number;
  keyword: Array<string>;
  like_users: Array<any>;
  meta: Meta;
  origin: number;
  state: number;
  tags: Array<any>;
  title: string;
  updated_at: string;
}
export interface ArticleDetailParams {
  id: number;
  type: number;
}
export interface LikeParams {
  id: string;
  user_id: string;
}

// 文章列表
export interface ArticlesParams {
  keyword: string;
  likes: string; // 是否是热门文章
  state: number; // 文章发布状态 => 0 草稿，1 已发布,'' 代表所有文章
  tag_id: string;
  category_id: string;
  pageNum: number;
  pageSize: number;
}

export interface ArticleList {
  category: string[];
  created_at: number;
  desc: string;
  img_url: string;
  meta: Meta;
  tags: string[];
  title: string;
  id: number;
}
export interface ArticlesData {
  count: number;
  list: ArticleList[] | any;
}


// 留言
export interface MessageParams {
  email: string;
  phone: string;
  name: string;
  content: string;
}
export interface RulesItem {
  validator: Function;
  trigger: string;
}
export interface Rules {
  email: RulesItem[];
  phone: RulesItem[];
  name: RulesItem[];
  content: RulesItem[];
}


// 项目
export interface ProjectList {
  content: string;
  end_time: string;
  img: string;
  start_time: string;
  title: string;
  url: string;
  id: number;
}
export interface ProjectsData {
  count: number;
  list: ProjectList | any;
}


// 历程
export interface TimelineList {
  content: string;
  end_time: string;
  start_time: string;
  title: string;
  state: number;
  id: number;
}
export interface TimelinesData {
  count: number;
  list: TimelineList | any;
}


// 标签
export interface TagList {
  name: string;
  id: string;
}
export interface TagsData {
  count: number;
  list: TagList | any;
}