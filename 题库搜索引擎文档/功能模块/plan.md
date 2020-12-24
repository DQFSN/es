# model

## 教案题目/PlanItem

- item_type：int	题目类型
- tag_ids：list     包括的知识点
- difficulty：int。难度
- item_id：ObjectID， 题目id

## EpPart /例题

- tag_id：ObjectID，知识点id

- items：list，对应的题目ID

## PlanDefaults/默认教案

- ep_classes
- ex_classes
- cmk_classes

## Plan/教案

- 数据库表名：plan
- name
- owner_id
- grade
- remark
- defaults：PlanDefaults
- ep_parts：lsit of EpPart
- ex_items：list of PlanItem
- cmt_items：list of int
- mtime
- ctime

## Plancollection

- 数据库表名：plan_collection

- name

- owner_id

- remark

- plan_ids：plan的id

- i_owner_id_1_id_1：含义不明

  ```python
  i_owner_id_1_id_1 = Index().ascending('owner_id').ascending('_id')
  ```

  

# handlers

和教案相关的操作

- URL

  ```json
  /plan_collections
  /plan_collection/<oid:id>
  /plans
  /plan_prepare
  /plan/<oid:plan_id>
  
  /latex/plan/<oid:plan_id>.pdf
  /latex/plan/<oid:plan_id>
  
  /api/plans
  /api/plan/<oid:plan_id>
  /api/plan
  /api/plan/<oid:plan_id>
  
  /api/plan_collection/<oid:id>
  /api/plan_collections
  ```

- export_plan_pdf

  - url
  - 作用