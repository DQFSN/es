# 题目搜索

## 1 、进入页面加载ktag

- GET /api/ktags

- 无请求参数

- response

  ```json
  {
  	"results": [{
  		"_id": "50e227000000000000001101",
  		"deleted": false,
  		"keywords": [],
  		"name": "载体",
  		"parent_id": null,
  		"stats": {
  			"avg_diff": 0.0,
  			"freq": 0.0
  		},
  		"type": 1101,
  		"weight": 0
  	}, {
  		"_id": "50e227000000000000001102",
  		"deleted": false,
  		"keywords": [],
  		"name": "知识点",
  		"parent_id": null,
  		"stats": {
  			"avg_diff": 2.8873854212134438,
  			"freq": 0.5153846153846153
  		},
  		"type": 1102,
  		"weight": 0
  	}, {
  		"_id": "50e227000000000000001103",
  		"deleted": false,
  		"keywords": [],
  		"name": "方法",
  		"parent_id": null,
  		"stats": {
  			"avg_diff": 3.158385093167702,
  			"freq": 0.4076923076923077
  		},
  		"type": 1103,
  		"weight": 0
  	}, {
  		"_id": "50e227000000000000001104",
  		"deleted": false,
  		"keywords": [],
  		"name": "专题",
  		"parent_id": null,
  		"stats": {
  			"avg_diff": 0.0,
  			"freq": 0.0
  		},
  		"type": 1104,
  		"weight": 0
  	}]
  ```

  

##  2、检索题目

- GET /api/item_search

- 参数

  - difficulty
  - keywords
  - edu=4
  - item_type=1000
  - classes
  - tag_ids
  - page_num=1
  - page_size=10

- response

  ```json
  {
  	"page_num": 1,
  	"page_size": 10,
  	"results": [{
  		"_id": "59093205060a05000970b2ab"
  	}, {
  		"_id": "599165c12bfec200011e002c"
  	}, {
  		"_id": "599165c12bfec200011e0064"
  	}, {
  		"_id": "59125bf7e020e7000878f6bb"
  	}, {
  		"_id": "599165c12bfec200011e0063"
  	}, {
  		"_id": "599165c12bfec200011e0062"
  	}, {
  		"_id": "599165c12bfec200011e002f"
  	}, {
  		"_id": "590c1439d42ca7000a7e7e43"
  	}, {
  		"_id": "599165c12bfec200011e002d"
  	}, {
  		"_id": "599165c12bfec200011e0066"
  	}],
  	"total": 21653
  }
  ```

  

## 3、请求题目by id（未完成联表查询）

- GET  /api/item/id

- 参数

  - render_kcode=1
  - &with_tagging=1
  - &paper_name_as_origin=1
  - &with_hydra=1

- response

  ```json
  {
  	"_id": "59093205060a05000970b2ab",
  	"boost": 1.0,
  	"circle": {
  		"_id": "5996c4966e9924000131c4ca",
  		"deleted": false,
  		"digest": "坐标平面上点集s满足s={(xy)∣\\log_2(x^2-x+2)=2\\sin^4y+2\\cos^4yy∈[-\\dfracπ8\\dfracπ4]}将点集s中所有点向x轴作投影所得投影线段长度之和",
  		"hydra_roles": [],
  		"item_edu": 4,
  		"item_ids": ["59093205060a05000970b2ab"],
  		"item_qs_size": 1,
  		"item_type": 1002,
  		"ngrams": ["标平", "坐标", "平面"],
  		"ngrams1": ["之和", "长度", "度之"]
  	},
  	"ctime": 1493774853.317623,
  	"data": {
  		"abilities": {
  			"abstract": 0,
  			"appl": 0,
  			"calc": 3,
  			"data": 0,
  			"reason": 3,
  			"spatial": 0
  		},
  		"categories": [2],
  		"difficulty": 3,
  		"ktags": [],
  		"qs": [{
  			"ans": "$2$",
  			"desc": "坐标平面上的点集&nbsp;$S$&nbsp;满足$$S=\\left\\{(x,y)\\mid&nbsp;{\\log_2}(x^2-x+2)=2\\sin^4y+2\\cos^4y,y\\in\\left[-\\dfrac{\\mathrm&nbsp;\\pi}&nbsp;8,\\dfrac{\\mathrm&nbsp;\\pi}&nbsp;4\\right]\\right\\},$$将点集&nbsp;$S$&nbsp;中所有点向&nbsp;$x$&nbsp;轴作投影，所得投影线段的长度之和为<nn>         </nn>．",
  			"exp": "由于&nbsp;$2\\sin^4y+2\\cos^4y=2-\\sin^22y$，于是&nbsp;$2\\sin^4y+2\\cos^4y$&nbsp;的取值范围是&nbsp;$\\left[1,2\\right]$，这样就可得到了$$1\\leqslant&nbsp;{\\log_2}(x^2-x+2)\\leqslant&nbsp;2,$$即$$\\dfrac&nbsp;14\\leqslant&nbsp;\\left(x-\\dfrac&nbsp;12\\right)^2\\leqslant&nbsp;\\dfrac&nbsp;94,$$因此所求的投影线段的长度之和为&nbsp;$2\\left(\\sqrt{\\dfrac&nbsp;94}-\\sqrt{\\dfrac&nbsp;14}\\right)=2$．",
  			"ktags": [{
  				"auto": false,
  				"id": "5901ddfc060a05000980aef3",
  				"path": [{
  					"_id": "50e227000000000000001102",
  					"name": "知识点",
  					"type": 1102
  				}, {
  					"_id": "5901d16a060a05000980aedb",
  					"name": "函数",
  					"type": 1102
  				}, {
  					"_id": "5901d365060a050008e6212a",
  					"name": "函数的图象与性质",
  					"type": 1102
  				}, {
  					"_id": "5901ddfc060a05000980aef3",
  					"name": "函数的最值和值域",
  					"type": 1102
  				}],
  				"primary": false
  			}, {
  				"auto": false,
  				"id": "5901d8a3060a05000bf29095",
  				"path": [{
  					"_id": "50e227000000000000001102",
  					"name": "知识点",
  					"type": 1102
  				}, {
  					"_id": "5901d16a060a05000980aedb",
  					"name": "函数",
  					"type": 1102
  				}, {
  					"_id": "5901d33f060a05000980aedf",
  					"name": "常见初等函数",
  					"type": 1102
  				}, {
  					"_id": "5901d8a3060a05000bf29095",
  					"name": "三角函数",
  					"type": 1102
  				}],
  				"primary": false
  			}, {
  				"auto": false,
  				"id": "5901d88c060a05000a4a9752",
  				"path": [{
  					"_id": "50e227000000000000001102",
  					"name": "知识点",
  					"type": 1102
  				}, {
  					"_id": "5901d16a060a05000980aedb",
  					"name": "函数",
  					"type": 1102
  				}, {
  					"_id": "5901d33f060a05000980aedf",
  					"name": "常见初等函数",
  					"type": 1102
  				}, {
  					"_id": "5901d88c060a05000a4a9752",
  					"name": "对数函数",
  					"type": 1102
  				}],
  				"primary": false
  			}],
  			"opts": [],
  			"remark": "每日一题[594]影子有多长．",
  			"root_item_id": "59093205060a05000970b2ab",
  			"step": ""
  		}],
  		"stem": "",
  		"subtype": 1,
  		"type": 1002
  	},
  	"digest": "坐标平面上点集s满足s={(xy)∣\\log_2(x^2-x+2)=2\\sin^4y+2\\cos^4yy∈[-\\dfracπ8\\dfracπ4]}将点集s中所有点向x轴作投影所得投影线段长度之和",
  	"extra": {
  		"classes": [7],
  		"edu": 4,
  		"origins": []
  	},
  	"hydra": null,
  	"lock_by_uid": null,
  	"lock_time": 0.0,
  	"mtime": 1509102896.6859298,
  	"simhash": "",
  	"state": 5,
  	"ver": 20
  }
  ```

  

# 知识点编辑

分为四部分，四部分的请求个数和方式完全相同

首先 /api/ktags+type参数，请求同一type下的所有ktag，然后请求这个type下的父节点（ID貌似是直接指定的）的信息，请求格式：/api/ktag/50e227000000000000001102，50e227000000000000001102为root节点ID。

## 节点增删改查：

- 增： PUT /api/ktags 	带json

- 删：PATCH   /api/ktag/节点ID	软删除，json中{deleted：true}

- 改：PATCH   /api/ktag/节点ID	json

- 查：GET		/api/ktag/节点ID	

## 1、知识点

 - 请求1 	

   - Url 

     GET   /api/ktags

   - 参数     

     type=1102

   - response

     ```json
     	"results": [{
     		"_id": "5c7f59f1210b28428f14d17c",
     		"deleted": false,
     		"keywords": [],
     		"name": "高中视角下的解析几何",
     		"parent_id": "50e227000000000000001102",
     		"stats": {
     			"avg_diff": 0.0,
     			"freq": 0.0
     		},
     		"type": 1102,
     		"weight": -51
     	}, {
     		"_id": "59253ef582e8bd0009968414",
     		"deleted": false,
     		"keywords": [],
     		"name": "利用函数不等式进行估值",
     		"parent_id": "5901e1cc060a05000a4a976a",
     		"stats": {
     			"avg_diff": 3.75,
     			"freq": 0.015384615384615385
     		},
     		"type": 1102,
     		"weight": -34
     	}]
     ```

     

- 请求2

  - URL

    GET 	 api/ktag/50e227000000000000001102

  - 参数

    无

  - response

    ```json
    {
    	"_id": "50e227000000000000001102",
    	"assess_dirs": [],
    	"deleted": false,
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"keywords": [],
    	"name": "知识点",
    	"parent_id": null,
    	"preconditions": [],
    	"stats": {
    		"avg_diff": 2.8873854212134438,
    		"freq": 0.5153846153846153
    	},
    	"teaching_objective": {
    		"desc": "",
    		"lesson_count": 0.0,
    		"level": ""
    	},
    	"type": 1102,
    	"weight": 0,
    	"wiki": ""
    }
    ```

- 修改知识点

  - URL

    PATCH		api/ktag/50e227000000000000001102

  - 参数

    ```json
    {
    	"name": "知识点",
    	"deleted": false,
    	"weight": 0,
    	"wiki": "",
    	"assess_dirs": [],
    	"teaching_objective": {
    		"desc": "",
    		"lesson_count": 0,
    		"level": "A"
    	},
    	"keywords": [],
    	"desc_tex": "",
    	"extra_desc_tex": ""
    }
    ```

    

  - Response

    ```json
    {
    	"_id": "50e227000000000000001102",
    	"assess_dirs": [],
    	"deleted": false,
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"keywords": [],
    	"name": "知识点",
    	"parent_id": null,
    	"preconditions": [],
    	"stats": {
    		"avg_diff": 2.8873854212134438,
    		"freq": 0.5153846153846153
    	},
    	"teaching_objective": {
    		"desc": "",
    		"lesson_count": 0.0,
    		"level": "A"
    	},
    	"type": 1102,
    	"weight": 0,
    	"wiki": ""
    }
    ```

- 增加子节点

  - URL

    PUT		api/ktag

  - 参数

    ```json
    {
    	"_id": null,
    	"name": "测试知识点的子节点",
    	"type": 1102,
    	"weight": 0,
    	"wiki": "单身的",
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"deleted": false,
    	"assess_dirs": [1, 2],
    	"path": "知识点 > 知识点的子节点",
    	"parent_id": "50e227000000000000001102",
    	"teaching_objective": {
    		"desc": "单身的",
    		"level": "A",
    		"lesson_count": 10
    	},
    	"stats": {
    		"freq": 0,
    		"avg_diff": 0
    	},
    	"keywords": ["测试"]
    }
    ```

    

  - response

    ```json
    {
    	"_id": "5fcf3081210b2863adb27360",
    	"assess_dirs": [1, 2],
    	"deleted": false,
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"keywords": ["测试"],
    	"name": "测试知识点的子节点",
    	"parent_id": "50e227000000000000001102",
    	"preconditions": [],
    	"stats": {
    		"avg_diff": 0.0,
    		"freq": 0.0
    	},
    	"teaching_objective": {
    		"desc": "单身的",
    		"lesson_count": 10.0,
    		"level": "A"
    	},
    	"type": 1102,
    	"weight": 0,
    	"wiki": "单身的"
    }
    ```

- 删除节点(软删除)

  - URL

    PATCH	api/ktag/5fcf3081210b2863adb27360

  - 参数

    ```json
    {
      "deleted":true
    }
    ```

    

  - response

    ```json
    {
    	"_id": "5fcf3081210b2863adb27360",
    	"assess_dirs": [1, 2],
    	"deleted": true,
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"keywords": ["测试"],
    	"name": "测试知识点的子节点",
    	"parent_id": "50e227000000000000001102",
    	"preconditions": [],
    	"stats": {
    		"avg_diff": 0.0,
    		"freq": 0.0
    	},
    	"teaching_objective": {
    		"desc": "单身的",
    		"lesson_count": 10.0,
    		"level": "A"
    	},
    	"type": 1102,
    	"weight": 0,
    	"wiki": "单身的"
    }
    ```

    

## 2、方法

 - 请求1 	

    - url 

      /api/ktags

   - 参数     

     type=1103

   - response

     ```json
     "results": [{
     		"_id": "5a13dc57aaa1af00079cad2e",
     		"deleted": false,
     		"keywords": [],
     		"name": "算法与程序框图",
     		"parent_id": "5902ff21060a05000a4a97ba",
     		"stats": {
     			"avg_diff": 0.0,
     			"freq": 0.0
     		},
     		"type": 1103,
     		"weight": -30
     	}, {
     		"_id": "5925440382e8bd0007792069",
     		"deleted": false,
     		"keywords": ["状态量"],
     		"name": "构造状态量",
     		"parent_id": "5902ff21060a05000a4a97ba",
     		"stats": {
     			"avg_diff": 3.3076923076923075,
     			"freq": 0.03461538461538462
     		},
     		"type": 1103,
     		"weight": -21
     	}, {
     		"_id": "592543e882e8bd000aa6acbc",
     		"deleted": false,
     		"keywords": ["半不变量"],
     		"name": "构造半不变量",
     		"parent_id": "5902ff21060a05000a4a97ba",
     		"stats": {
     			"avg_diff": 4.0,
     			"freq": 0.023076923076923078
     		},
     		"type": 1103,
     		"weight": -20
     	}]
     ```

     

- 请求2

   - url

     api/ktag/50e227000000000000001103

  - 参数

    无

  - response

    ```json
    {
    	"_id": "50e227000000000000001102",
    	"assess_dirs": [],
    	"deleted": false,
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"keywords": [],
    	"name": "知识点",
    	"parent_id": null,
    	"preconditions": [],
    	"stats": {
    		"avg_diff": 2.8873854212134438,
    		"freq": 0.5153846153846153
    	},
    	"teaching_objective": {
    		"desc": "",
    		"lesson_count": 0.0,
    		"level": ""
    	},
    	"type": 1102,
    	"weight": 0,
    	"wiki": ""
    }
    ```

- 增删改同知识点，更换请求URL中的ID即可

## 3、题型

 - 请求1 	

    - url 

      /api/ktags

   - 参数     

     type=1201

   - response

     ```json
     {
     	"results": [{
     		"_id": "5903010a060a05000a4a97c1",
     		"deleted": false,
     		"keywords": [],
     		"name": "解析几何创新题",
     		"parent_id": "5901da0c060a05000bf29099",
     		"stats": {
     			"avg_diff": 3.2142857142857144,
     			"freq": 0.03461538461538462
     		},
     		"type": 1201,
     		"weight": -60
     	}, {
     		"_id": "59200216623a97000a198dc8",
     		"deleted": false,
     		"keywords": [],
     		"name": "圆锥曲线的性质证明问题",
     		"parent_id": "5901da0c060a05000bf29099",
     		"stats": {
     			"avg_diff": 3.4423076923076925,
     			"freq": 0.07692307692307693
     		},
     		"type": 1201,
     		"weight": -40
     	}, {
     		"_id": "59030102060a05000980af5b",
     		"deleted": false,
     		"keywords": [],
     		"name": "概率创新题",
     		"parent_id": "5901d9ff060a05000a4a9756",
     		"stats": {
     			"avg_diff": 3.0,
     			"freq": 0.0038461538461538464
     		},
     		"type": 1201,
     		"weight": -30
     	}]
     }
     ```

     

- 请求2

   - url

     api/ktag/50e227000000000000001201

  - 参数

    无

  - response

    ```json
    {
    	"_id": "50e227000000000000001201",
    	"assess_dirs": [],
    	"deleted": false,
    	"desc": "",
    	"desc_tex": "",
    	"extra_desc": "",
    	"extra_desc_tex": "",
    	"keywords": [],
    	"name": "题型",
    	"parent_id": null,
    	"preconditions": [],
    	"stats": {
    		"avg_diff": 3.1469949312092687,
    		"freq": 0.5
    	},
    	"teaching_objective": {
    		"desc": "",
    		"lesson_count": 0.0,
    		"level": ""
    	},
    	"type": 1201,
    	"weight": 0,
    	"wiki": ""
    }
    ```

    

## 4、载体

 - 请求1 	

    - url 

      /api/ktags

   - 参数     

     type=1101

   - response

     ```json
     {
     	"_id": "50e227000000000000001101",
     	"assess_dirs": [],
     	"deleted": false,
     	"desc": "",
     	"desc_tex": "",
     	"extra_desc": "",
     	"extra_desc_tex": "",
     	"keywords": [],
     	"name": "载体",
     	"parent_id": null,
     	"preconditions": [],
     	"stats": {
     		"avg_diff": 0.0,
     		"freq": 0.0
     	},
     	"teaching_objective": {
     		"desc": "",
     		"lesson_count": 0.0,
     		"level": ""
     	},
     	"type": 1101,
     	"weight": 0,
     	"wiki": ""
     }
     ```

     

- 请求2

  无

# 教案箱

- 教案和试卷是同一张表的不同呈现形式？

## 1、教案箱

 - URL

   /api/plans

- 参数
  - owner_id=5fb33927210b2863adb2734d
  - grade=0
  - keywords=
  - remark=

- response

  ```json
  {
  	"page_no": 1,
  	"page_size": 20,
  	"results": [{
  		"_id": "5fc72ba6210b2863acf5ae0f",
  		"cmk_items": [],
  		"ctime": 1606888358.2615147,
  		"defaults": {
  			"cmk_classes": [],
  			"ep_classes": [],
  			"ex_classes": []
  		},
  		"ep_parts": [{
  			"items": [{
  				"difficulty": 0,
  				"item_id": "5966ebf2030398000bbee7ef",
  				"item_type": 1001,
  				"tag_ids": ["598052a9400acd00094aaa7c"]
  			}],
  			"tag_id": "598052a9400acd00094aaa7c"
  		}, {
  			"items": [{
  				"difficulty": 0,
  				"item_id": "599165c22bfec200011e02ff",
  				"item_type": 1002,
  				"tag_ids": ["59266e83ee79c2000a59dc06"]
  			}],
  			"tag_id": "59266e83ee79c2000a59dc06"
  		}, {
  			"items": [{
  				"difficulty": 0,
  				"item_id": "59fbbc5e03bdb1000a37cc38",
  				"item_type": 1002,
  				"tag_ids": ["59811302400acd0007dcc3fd"]
  			}],
  			"tag_id": "59811302400acd0007dcc3fd"
  		}, {
  			"items": [{
  				"difficulty": 0,
  				"item_id": "5cb3ec86210b28021fc75550",
  				"item_type": 1001,
  				"tag_ids": ["5ea632db210b280d37822c7c"]
  			}, {
  				"difficulty": 0,
  				"item_id": "59111d4740fdc700073df559",
  				"item_type": 1001,
  				"tag_ids": ["5ea632db210b280d37822c7c"]
  			}],
  			"tag_id": "5ea632db210b280d37822c7c"
  		}],
  		"ex_items": [{
  			"difficulty": 0,
  			"item_id": "599165c72bfec200011e1364",
  			"item_type": 1002,
  			"tag_ids": ["598052a9400acd00094aaa7c"]
  		}, {
  			"difficulty": 0,
  			"item_id": "59fad8786ee16400083d2843",
  			"item_type": 1001,
  			"tag_ids": ["59266e83ee79c2000a59dc06"]
  		}],
  		"grade": 43,
  		"mtime": 1606889209.9038281,
  		"name": "集合与映射",
  		"owner_id": "5d19c082210b28021fc77bda",
  		"remark": ""
  	}]
  }
  ```

  

## 2、合集

 - url

   api/plan_collections

 - 参数

   - name=
   - owner_id=5fb33927210b2863adb2734d

- response

  ```json
  {
  	"page_no": 1,
  	"page_size": 20,
  	"results": [],
  	"total": 0
  }
  ```

  

# 我的试卷|题集

## 1、试卷

- 查询
  - URL

    api/papers

  - 参数

    - owner_id=
    - grade=0
    - keywords=
    - remark=

  - response

    ```json
    "page_no": 1,
    	"page_size": 20,
    	"results": [{
    		"_id": "5faca394210b2863adb27348",
    		"ctime": 1605149588.1823654,
    		"defaults": {
    			"classes": []
    		},
    		"grade": 41,
    		"mtime": 1605149588.1823514,
    		"name": "1",
    		"owner_id": "5f193d46210b28774f713a0d",
    		"parts": [{
    			"item_type": 1001,
    			"items": [{
    				"difficulty": 0,
    				"item_id": "599165b62bfec200011de19d",
    				"item_type": 1001,
    				"tag_ids": ["5efaf427210b28017b0e2fbc"]
    			}, {
    				"difficulty": 0,
    				"item_id": "599165bf2bfec200011dfa74",
    				"item_type": 1001,
    				"tag_ids": ["5efaf427210b28017b0e2fbc"]
    			}]
    ```

- 增加

  - URL

    PUT		api/paper?

  - 参数

    ```json
    {
    	"name": "测试试卷添加",
    	"ktags": [{
    		"_id": "5efaf427210b28017b0e2fbc",
    		"diff": 3,
    		"parts": [{
    			"count": 3,
    			"type": 1001
    		}, {
    			"count": 2,
    			"type": 1003
    		}]
    	}, {
    		"_id": "59030077060a05000bf29103",
    		"diff": 0,
    		"parts": [{
    			"count": 3,
    			"type": 1001
    		}, {
    			"count": 2,
    			"type": 1002
    		}, {
    			"count": 2,
    			"type": 1003
    		}]
    	}, {
    		"_id": "59367ad5c2b4e70008d3b94e",
    		"diff": 0,
    		"parts": [{
    			"count": 3,
    			"type": 1001
    		}, {
    			"count": 3,
    			"type": 1002
    		}]
    	}],
    	"grade": 41,
    	"defaults": {
    		"classes": []
    	}
    }
    ```

    

  - response

    ```json
    {
    	"_id": "5fd0b901210b2863acf5ae1a",
    	"ctime": 1607514369.347564,
    	"defaults": {
    		"classes": []
    	},
    	"grade": 41,
    	"mtime": 1607514369.3475363,
    	"name": "测试试卷添加",
    	"owner_id": "5fb33927210b2863adb2734d",
    	"parts": [{
    		"item_type": 1001,
    		"items": [{
    			"difficulty": 0,
    			"item_id": "59094c15060a05000970b36b",
    			"item_type": 1001,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 0,
    			"item_id": "59096fd839f91d0009d4bf96",
    			"item_type": 1001,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 0,
    			"item_id": "590bdc296cddca000861100e",
    			"item_type": 1001,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 0,
    			"item_id": "5912a422e020e7000a798bdd",
    			"item_type": 1001,
    			"tag_ids": ["59367ad5c2b4e70008d3b94e"]
    		}, {
    			"difficulty": 0,
    			"item_id": "59b9dfdcb3e1920008f9698b",
    			"item_type": 1001,
    			"tag_ids": ["59367ad5c2b4e70008d3b94e"]
    		}, {
    			"difficulty": 0,
    			"item_id": "5a0e7de8aaa1af00079ca9e2",
    			"item_type": 1001,
    			"tag_ids": ["59367ad5c2b4e70008d3b94e"]
    		}, {
    			"difficulty": 3,
    			"item_id": "599165b62bfec200011de19d",
    			"item_type": 1001,
    			"tag_ids": ["5efaf427210b28017b0e2fbc"]
    		}, {
    			"difficulty": 3,
    			"item_id": "599165ca2bfec200011e1b36",
    			"item_type": 1001,
    			"tag_ids": ["5efaf427210b28017b0e2fbc"]
    		}]
    	}, {
    		"item_type": 1002,
    		"items": [{
    			"difficulty": 0,
    			"item_id": "59101c7d857b4200092b0811",
    			"item_type": 1002,
    			"tag_ids": ["59367ad5c2b4e70008d3b94e"]
    		}, {
    			"difficulty": 0,
    			"item_id": "593539f5c2b4e7000a0853dc",
    			"item_type": 1002,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 0,
    			"item_id": "59872e675ed01a0008fa5ef1",
    			"item_type": 1002,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 0,
    			"item_id": "59c71435778d4700085f6bbd",
    			"item_type": 1002,
    			"tag_ids": ["59367ad5c2b4e70008d3b94e"]
    		}, {
    			"difficulty": 0,
    			"item_id": "5a01488f03bdb1000a37d0ea",
    			"item_type": 1002,
    			"tag_ids": ["59367ad5c2b4e70008d3b94e"]
    		}]
    	}, {
    		"item_type": 1003,
    		"items": [{
    			"difficulty": 0,
    			"item_id": "59082029060a050008e621e1",
    			"item_type": 1003,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 0,
    			"item_id": "5959d640d3b4f9000ad5ea3e",
    			"item_type": 1003,
    			"tag_ids": ["59030077060a05000bf29103"]
    		}, {
    			"difficulty": 3,
    			"item_id": "5efe9a34210b28017ae2fdd4",
    			"item_type": 1003,
    			"tag_ids": ["5efaf427210b28017b0e2fbc"]
    		}]
    	}],
    	"remark": ""
    }
    ```

    

## 2、我的题集

 - url 

   api/omega_papers

- 参数

  - name=
  - year=0
  - grade=0
  - cat=
  - uid=5fb33927210b2863adb2734d

- response

  ```json
  {
  	"page_no": 1,
  	"page_size": 20,
  	"results": [{
  		"_id": "5fcde477210b2863adb27359",
  		"cat": "备选题",
  		"deleted": false,
  		"districts": ["四川"],
  		"doc": "",
  		"edu": 4,
  		"flow_stats": {
  			"no_exp": 0,
  			"review1": 1,
  			"review2": 0,
  			"start": 0,
  			"typeset": 0
  		},
  		"grade": 43,
  		"item_ids": ["599165c22bfec200011e0581"],
  		"item_tpls": {},
  		"name": "test",
  		"remark": "",
  		"uid": "5fb33927210b2863adb2734d",
  		"year": 2020
  	}],
  	"total": 1
  }
  ```

  

## 3、 真题试卷

 - URL

   api/omega_papers

- 参数
  - name=
  - year=0
  - grade=0
  - cat=
  - uid=

- response

  ```json
  {
  	"page_no": 1,
  	"page_size": 20,
  	"results": [{
  		"_id": "5f264451210b2865a67885fe",
  		"cat": "自招真题",
  		"deleted": false,
  		"districts": [],
  		"doc": "",
  		"edu": 4,
  		"flow_stats": {
  			"no_exp": 3,
  			"review1": 0,
  			"review2": 0,
  			"start": 0,
  			"typeset": 12
  		},
  		"grade": 43,
  		"item_ids": ["5f264523210b2865a6788600", "5f2646ba210b2865a6788607", "5f264730210b2865a678860d", "5f264867210b2865a8643c27", "5f2648cf210b2865a6788614", "5f264ec1210b2865a678861b", "5f265a40210b2865a8643c2f", "5f265aba210b2865a6788624", "5f265b7e210b2865a678862c", "5f265bee210b2865a8643c3c", "5f265c54210b2865a6788635", "5f265cfe210b2865a8643c42"],
  		"item_tpls": {},
  		"name": "2020年北大强基考试数学回忆",
  		"remark": "",
  		"uid": null,
  		"year": 2020
  	}, {
  		"_id": "5f1931e9210b28775079b891",
  		"cat": "高考真题",
  		"deleted": false,
  		"districts": ["北京"],
  		"doc": "",
  		"edu": 4,
  		"flow_stats": {
  			"no_exp": 111,
  			"review1": 12,
  			"review2": 0,
  			"start": 0,
  			"typeset": 112
  		},
  		"grade": 43,
  		"item_ids": ["5f19328b210b28774f7139c9", "5f1933f1210b28774f7139d9", "5f1935ba210b28775079b8aa", "5f1936ab210b28774f7139e7", "5f1937a5210b28775079b8c1", "5f193990210b28774f7139f8", "5f193da1210b28775079b8e4", "5f194073210b28775079b8f3", "5f19439c210b28775079b901", "5f194e75210b28775079b90e", "5f194f0c210b28775079b918", "5f1a5138210b2865a6788342", "5f1a649d210b2865a678834c", "5f1a658e210b2865a6788353", "5f1a661d210b2865a678835b", "5f1a6707210b2865a8643a3e", "5f1a676a210b2865a6788364", "5f1a6859210b2865a678836b", "5f1a691b210b2865a6788373", "5f1a6962210b2865a8643a47", "5f1a69aa210b2865a8643a4c", "5f1a69ed210b2865a8643a52", "5f1a6a36210b2865a6788383", "5f1a6a9a210b2865a8643a58", "5f1a7abe210b2865a8643a5e", "5f1a7b22210b2865a678838f", "5f1a7e97210b2865a6788398", "5f1a7f31210b2865a8643a68", "5f1a8060210b2865a8643a6d", "5f1a8260210b2865a67883a5", "5f1a8317210b2865a8643a78", "5f1a83fc210b2865a67883ae", "5f1a8470210b2865a67883b5", "5f1a8539210b2865a8643a82", "5f1a8d8b210b2865a8643a8a", "5f1a8e12210b2865a67883c1", "5f1a8e80210b2865a8643a90", "5f1a8ecb210b2865a67883cb", "5f1a90c8210b2865a8643a99", "5f1a9141210b2865a8643a9f", "5f1a933d210b2865a67883d5", "5f1a9492210b2865a8643aa6", "5f1a94f2210b2865a8643aac", "5f1a9549210b2865a8643ab2", "5f1a9725210b2865a8643ab9", "5f1a9780210b2865a67883e6", "5f1a987a210b2865a8643ac1", "5f1a9901210b2865a67883f0", "5f1a99b9210b2865a67883fe", "5f1a9a3a210b2865a8643ac7", "5f1a9aaf210b2865a6788407", "5f1a9b69210b2865a6788410", "5f1a9bfa210b2865a6788419", "5f1a9c4c210b2865a678841f", "5f1a9e09210b2865a678842c", "5f1a9ea4210b2865a6788435", "5f1a9f58210b2865a8643adb", "5f1aa039210b2865a8643ae1", "5f1aa0b9210b2865a8643ae8", "5f1aa106210b2865a8643aee", "5f1aa346210b2865a6788446", "5f1aa399210b2865a8643afa", "5f1aa3e4210b2865a678844f", "5f1aa42a210b2865a6788456", "5f1aa4d9210b2865a678845c", "5f1aa5d0210b2865a6788464", "5f1aa642210b2865a8643b08", "5f1aa693210b2865a678846b", "5f1aa704210b2865a6788471", "5f1aa786210b2865a8643b13", "5f1aa84a210b2865a8643b18", "5f1aaa99210b2865a678847d", "5f1aab60210b2865a8643b21", "5f1aac49210b2865a6788483", "5f1aacc2210b2865a8643b2a", "5f1aad85210b2865a8643b30", "5f1ab6da210b2865a678849d", "5f1ab768210b2865a67884a3", "5f1ab7e2210b2865a67884a9", "5f1ab91b210b2865a67884b2", "5f1ab9c2210b2865a67884ba", "5f1aba09210b2865a67884c0", "5f1b945e210b2865a67884ca", "5f1b95d4210b2865a67884d2", "5f1b9ae1210b2865a8643b49", "5f1b9b35210b2865a67884db", "5f1b9b9f210b2865a8643b4f", "5f1b9be9210b2865a67884e4", "5f1b9c42210b2865a67884ea", "5f1b9cf6210b2865a67884f1", "5f1b9dd4210b2865a67884f8", "5f1b9e3d210b2865a67884fe", "5f1b9f21210b2865a8643b5f", "5f1ba07b210b2865a6788505", "5f1ba438210b2865a678850a", "5f1ba48d210b2865a8643b69", "5f1ba534210b2865a6788512", "5f1ba56f210b2865a6788518", "5f1ba5c0210b2865a6788520", "5f1bae68210b2865a8643b75", "5f1bb56c210b2865a6788528", "5f1bb6e9210b2865a678852f", "5f1bbbe0210b2865a8643b7d", "5f1bbc48210b2865a8643b83", "5f1bbcd1210b2865a6788538", "5f1bbd27210b2865a678853e", "5f1bbd62210b2865a8643b8d", "5f1bc20a210b2865a6788545", "5f1bc27f210b2865a8643b97", "5f1bc35d210b2865a678854e", "5f1bc3ae210b2865a8643b9c", "5f1bc61d210b2865a678855b", "5f1bc6f8210b2865a6788562", "5f1bc759210b2865a6788568", "5f1bc7c6210b2865a8643ba7", "5f1bc875210b2865a8643bb0", "5f1bc8d0210b2865a6788571", "5f1bc8f1210b2865a8643bb6", "5f1bc92c210b2865a6788579", "5f1bc9d6210b2865a8643bbd", "5f1bca16210b2865a6788584", "5f1bca86210b2865a8643bc2", "5f1bcacb210b2865a678858f", "5f1bcb33210b2865a8643bc8"],
  		"item_tpls": {},
  		"name": "五年高考分类汇编（北京版）补充题",
  		"remark": "",
  		"uid": null,
  		"year": 2020
  	}, {
  		"_id": "5f0c020e210b28775079b1d6",
  		"cat": "高考真题",
  		"deleted": false,
  		"districts": ["海南"],
  		"doc": "",
  		"edu": 4,
  		"flow_stats": {
  			"no_exp": 2,
  			"review1": 22,
  			"review2": 0,
  			"start": 0,
  			"typeset": 0
  		},
  		"grade": 43,
  		"item_ids": ["5f0c0223210b28774f71358c", "5f0c0225210b28775079b1d8", "5f0c0227210b28775079b1dd", "5f0c0229210b28774f713594", "5f0c022b210b28775079b1e2", "5f0c022d210b28775079b1e8", "5f0c022f210b28774f71359b", "5f0c0231210b28775079b1ee", "5f0c05e9210b28775079b205", "5f0c05ec210b28774f7135b5", "5f0c05ed210b28775079b20b", "5f0c05f0210b28775079b211", "5f0c0890210b28775079b226", "5f0c0891210b28774f7135c6", "5f0c0892210b28775079b22c", "5f0c0893210b28774f7135cc", "5f0d607d210b28774f713661", "5f0d60b5210b28775079b31d", "5f0d616d210b28775079b325", "5f0d61f0210b28774f71366c", "5f0d628d210b28775079b335", "5f0d634b210b28775079b33c"],
  		"item_tpls": {},
  		"name": "2020年新高考（Ⅱ）卷",
  		"remark": "",
  		"uid": null,
  		"year": 2020
  	}]
  }
  ```

  