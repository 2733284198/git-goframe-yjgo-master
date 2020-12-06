/*
==========================================================================
云捷GO自动生成菜单SQL,只生成一次,按需修改.
生成日期：2020-12-07 01:38:53
生成路径: document/sql/module/zjuser_menu.sql
生成人：yunjie
==========================================================================
*/

-- 菜单 SQL
insert into sys_menu (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('专家用户', '4', '1', '/module/zjuser', 'C', '0', 'zjuser:view', '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '专家用户菜单');

-- 按钮父菜单ID
SELECT @parentId := LAST_INSERT_ID();

-- 按钮 SQL
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('专家用户查询', @parentId, '1',  '#',  'F', '0', 'module:zjuser:list',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('专家用户新增', @parentId, '2',  '#',  'F', '0', 'module:zjuser:add',          '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('专家用户修改', @parentId, '3',  '#',  'F', '0', 'module:zjuser:edit',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('专家用户删除', @parentId, '4',  '#',  'F', '0', 'module:zjuser:remove',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('专家用户导出', @parentId, '5',  '#',  'F', '0', 'module:zjuser:export',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');