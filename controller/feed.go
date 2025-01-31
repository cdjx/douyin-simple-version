package controller

import (
	"douyin-simple-version/function"
	"douyin-simple-version/public"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	public.Response
	VideoList []public.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")

	if user, exits := usersLoginInfo[token]; exits {
		feeds, status := function.Query_feeds(user, c)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  status,
			VideoList: feeds,
			NextTime:  time.Now().Unix(),
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: public.Response{StatusCode: 1, StatusMsg: "User doesn't login"},
		})
	}
}

/*
Vedio Test URL

[
    {
        "attribution": "Liyao Xie / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdM.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcgc.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhXT.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/a8c412fa-f696-4ff2-9c76-e8ed9cdffe0f/604a87fc-e7bc-463e-8d56-cde7e661d690.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/ba258271-89c7-47bc-9742-bcae67c23202/f7ff4fe4-1346-47bb-9466-3f4662c1ac3a.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/b7014b7e-b38f-4a64-bd95-4a28a8ef6dee/113a2bf3-3a5f-45d4-8b6f-e40ce8559da3.mp4"
        }
    },
    {
        "attribution": "Yiming Li / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRG.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe5M.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiHa.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/0b927d99-e38a-4f51-8d1a-598fd4d6ee97/3493c85c-f35a-488f-9a8f-633e747fb141.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/bc3e9341-3243-4d2c-8469-940fef56ca2d/4720a02b-eabd-4593-a1d9-5c5d61916853.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/35960fe4-724f-44fc-ad77-0b91c55195e4/bfd49cd7-a0c6-467e-ae34-8674779e689b.mp4"
        }
    },
    {
        "attribution": "Schroptschop / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE58C.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa6N.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiHi.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/178161a4-26a5-4f84-96d3-6acea1909a06/2213bcd0-7d15-4da0-a619-e32d522572c0.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/b701c37f-3464-4d0a-a165-4a9468080ebd/7afe0873-a1f4-4fad-b771-a917687fcfc6.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/3d6f4af0-79ab-46fe-9d33-e191be5a878e/b4fa3f3e-a582-4bb5-9115-a82652e45b65.mp4"
        }
    },
    {
        "attribution": "Yiming Li / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE58B.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcge.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhXX.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/fe13f13c-c2cc-4998-b525-038b23bfa9b5/1a9d30ca-54be-411e-8b09-d72ef4488e05.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/6f18a83a-8d2a-4d01-b007-64fa411ea296/b619cc27-1c33-4d13-8488-d025aeaa7a8c.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/e6f8c1b2-b1ac-4343-a8bf-bd79a4d25380/9de24622-e13e-4741-95a6-04fdc39eb2b0.mp4"
        }
    },
    {
        "attribution": "Schroptschop / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe2I.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFq.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEab5.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/fb194c01-2ff6-4b4e-afbd-a00289124c4c/af7a74f5-5cb6-423e-b428-d05c0d36478d.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/66dcf148-272c-408c-82fc-b80bc8d0d533/be559c34-a832-4fc4-9573-02e896ba9c0b.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/da8b8a96-5a73-4a17-8ed8-960b8cd5afb1/30223688-167c-48a2-aa96-9764378a92d4.mp4"
        }
    },
    {
        "attribution": "Onfokus / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe2H.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFp.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiHj.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/e908e91f-370f-49ad-b4ce-775b7e7a05b4/a6287f74-46f0-42f9-b5d9-997f00585696.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/7d869c98-43ba-4e12-b789-de46ef2156a1/46aa0a69-7192-48b3-a96d-efee1a0bd4c3.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/c0f610f9-4fb1-4222-9ab5-a62f630a03ac/8caf2d97-2c57-48c9-8582-808b42089e3f.mp4"
        }
    },
    {
        "attribution": "Onfokus / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdO.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhUz.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2ef.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/15fc0eea-1091-4a28-a9b6-8caaf6013cf1/62b4c40f-ad3c-4099-a59a-bfbe324a461c.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/01fa79c3-2405-4415-bd6e-1d3b4e5efb29/c6c1e387-29fe-4062-b0ba-eea47a965bfb.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/45101561-a78f-4a5d-88f9-0efc175a1ae6/481ec582-3cae-4b68-8e2d-5833bcd2a593.mp4"
        }
    },
    {
        "attribution": "onuroner / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiEk.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhUA.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe9h.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/9752d732-2354-483f-a678-a6d0cd2c22b7/1a5ed13a-43f5-4e85-95c8-6579065c4d7c.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/13d2da4a-b642-4dc5-96ea-c13210066046/3fc2b744-a5f4-4031-be45-1e4c72010eb6.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/67699b85-15a5-4411-8e50-c3232b468f4d/0c4d59de-a4fc-4e82-86b7-3b0610febcdb.mp4"
        }
    },
    {
        "attribution": "Aflo Motion / Nimia ",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRH.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa6Q.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhXS.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/fb5e2f79-b18b-498a-98cf-d00352200957/50d763c7-96e6-492b-81ee-c692520504ef.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/df985854-2f4b-49b7-8046-4aca77755ecf/dbcbe54b-b368-4bbf-9e0c-ebbf58f6614c.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/a3ed4aeb-19cc-4dea-8c37-4e68d735bb43/dafa4eb0-f2e1-4c3c-84e9-d95eba8beb23.mp4"
        }
    },
    {
        "attribution": "Gavin Heffernan / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRI.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE5ak.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiHb.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/68f0e528-68ae-4180-9f0a-8536b7e10d6c/0e21b937-49a9-4e37-aa36-9bf7ae4a1983.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/ebaaf8bb-92d3-4d15-a9cd-338a2066f53d/1c6c0baf-e502-459b-bae2-d63876f80e1c.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/9dd41833-6a62-41c9-83c9-dec48c2fa6be/de7c1661-8429-4db6-9338-cf1a5046f660.mp4"
        }
    },
    {
        "attribution": "Banana Republic images / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdP.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElcv.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2ea.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/f9e318a4-7575-4809-b371-7b34a5807857/1f80c234-46af-41a5-8f50-8f0926e29a2c.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/ba83060f-19ec-45e8-919e-1f3f9bce7bc3/259b6592-7ef3-4b4e-ac7c-bb2abff6873b.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/24daadfe-8695-476e-8756-8338dcbeeb38/78c3a26d-1879-4dcf-b553-c3f7fc16c8ee.mp4"
        }
    },
    {
        "attribution": "anyaivanova / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRJ.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe5P.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEaaZ.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/87d3d62a-cfe2-4d03-8396-a1daf86b84ae/c2b325a7-0d9f-4b3f-a503-0568b15275b9.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/e891af95-67cc-43cf-b241-76dad145568f/b8b2ee7c-4d00-47e9-ac29-6abb54dd07ed.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/5e984120-bfbe-4e71-b422-689b4d490b52/1e04a945-b374-4920-b9bd-630d7f5683b5.mp4"
        }
    },
    {
        "attribution": "rusm / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa3q.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa6T.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElec.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/a75a7d73-21ab-4ac9-8c30-890433965c24/e9f6bdcb-eba0-4eca-b9d2-60d3415bf65f.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/0ad3f7e6-810d-468d-a3c3-55f4b1ab1fb5/542b2755-6599-4739-8e02-ae30a01e7905.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/fdf14e5b-06d7-44c4-ab56-751a5dbdaa14/947cfb92-b6ce-4ae8-b0b1-0e7bae2dd50b.mp4"
        }
    },
    {
        "attribution": "Yevgen Belich / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2aB.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE5al.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe9i.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/c66f076e-0032-4c28-bab5-51d28f650b4b/b09c5e47-dd44-4ee8-ad2c-6fadf08c66c0.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/e28aa34e-76df-4c9d-8083-90e03e6dc562/619c746b-f524-4f60-a70d-41d131fe0a00.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/910e7354-d9fa-46b7-851d-0dfbd7854b52/6b480c47-1959-400e-9fb5-22d51f4a6427.mp4"
        }
    },
    {
        "attribution": "rusm / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElaW.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2c6.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEled.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/e1c45064-728f-4251-be31-778fefb4e784/01ae5076-482e-421c-80ac-d1d9892e0a55.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/0a0fbcc9-ba13-4756-a69c-283160377b2e/8356cb12-6847-4e37-b165-35c525b6b405.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/9d7890f2-2646-4544-bbc7-5082440f52ef/df8ea196-566a-4291-85b2-c3bcb3781080.mp4"
        }
    },
    {
        "attribution": "Vladimir Solovyev / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElaX.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcgf.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2eb.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/394e73b6-e7ac-4b5d-8063-80d421afb446/0e09af42-f217-4deb-8580-fb0ab5420f6a.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/4b19dc45-879f-41df-aa0d-5ec8142b0962/e48d3d3f-d29b-4a66-ab78-d7c38e269435.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/5c5f1f0e-58c5-40f6-ac18-b5cfdf5182ba/9472a001-feeb-43c0-8506-c0bf10eceafe.mp4"
        }
    },
    {
        "attribution": "Qmin_CHIU / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElaY.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFs.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe9j.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/6bfef08a-9a8b-45b1-9c4b-338085feef4d/ee9aa021-c92f-403b-be35-0c7e5978bb07.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/0a760fce-8ad2-46ef-9165-9ccf2fc0c6ca/90146115-0c2f-44e7-bcdb-964faf55df0f.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/c2a00770-e5fb-4856-839b-92dbef440699/657d65c9-6890-4c09-bd44-b1f0e142b16b.mp4"
        }
    },
    {
        "attribution": "GOD_Cat / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe2J.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhUC.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiHf.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/ba33094e-efb2-480f-ad77-1a103af156d2/3b215666-f0b7-4538-9ea2-0637d5ed2ea3.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/2787cb8a-1de3-455d-82db-1b2ebfc6a2ba/53af17b5-5d5d-4a5d-bdb5-7fe03aaaeb79.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/b74edb16-445f-4ebf-9aa1-3f3744ed97b0/b7a5242c-05d3-4afd-8f36-23008f63b2b9.mp4"
        }
    },
    {
        "attribution": "Buddee Wiangngorn / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElaZ.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe5R.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiHe.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/b9614d22-e2ed-4689-becb-265d3d831c47/608d6881-b9a5-4ea9-aa8d-165bae3f43ec.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/0fe466af-1c3d-48d5-99d2-d4f869459d04/ded767be-be87-4351-b5a6-87b076c28101.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/12fe1894-69ae-4ae3-9d66-cb677d54bacf/67b9a312-27a5-4c1d-81ac-99730f54f4fc.mp4"
        }
    },
    {
        "attribution": "Jon Ingall / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiEl.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFt.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe9k.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/278b1c4f-3545-422c-9d56-242f18adadc9/2e3d6097-b233-40af-88f4-a90a74a33fa0.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/7995d65b-dafe-4996-8615-d30f6d740270/3bb3a5b7-396a-4d63-99f6-924810259974.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/41be5bb3-9d17-46f8-8540-d71c55be9cf5/27ddfed5-459a-4ecc-866b-62ba893b11fb.mp4"
        }
    },
    {
        "attribution": "si jiajun / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiEm.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa6V.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE5cU.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/9957a55f-0cb7-433c-b4fb-19a3b772ba37/5f0c70b2-ff2f-4c7d-8d05-03ae4553a930.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/131cf351-ec2c-4ade-96f1-387a2d3f8802/3fc3410c-db3a-45fb-bad3-b5fd8bbb21c6.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/d82eae7e-314b-4b14-9f9c-e6e00a9f20d8/5c7904f0-221e-4b82-ab90-1091786234e9.mp4"
        }
    },
    {
        "attribution": "d1sk / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiEn.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa6W.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE5cT.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/3f6ca508-598d-4cbd-a3e2-43deea7bc377/b60c553e-9f3f-4164-8850-700a9a73a899.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/907ab58a-8816-425c-bb84-6dcd9d6a643b/e200a3e4-40b7-4335-899f-04a62f5a8bd4.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/9f475b53-a3a8-40ff-882b-f392cee13609/58b01db3-4af9-440f-84ed-69dac2bfdf6d.mp4"
        }
    },
    {
        "attribution": "KropStock / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE58E.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElcy.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcjx.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/c633cd32-92ef-447f-876c-54c5c1a57484/d5a4b96f-01b4-4519-8470-62484e0b1b5b.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/0df8d2af-8070-45ba-b00d-e85e7cb17ad4/9c67d262-f53e-4358-9f99-beef9538e579.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/4f90a489-2c9e-4fa4-856a-c2b916489620/e1d3f80d-ec0b-4f2c-b13c-65e807b3ef24.mp4"
        }
    },
    {
        "attribution": "cherezoff / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa3v.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe5S.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcjw.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/b51e0ffe-ff7a-48a1-8599-805bf90e95e9/5b67fae0-b0ee-4b37-94aa-0b2ac26385ed.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/6e85cce0-b2fe-4c75-afca-d81ae0eed0bd/26df1714-e852-4644-93c6-a0ebf69c8bf7.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/16de71a0-ea6a-4918-b5b1-569ee4075eb7/afd3c9e6-de59-465c-bafa-250e9b46b29d.mp4"
        }
    },
    {
        "attribution": "BlackBoxGuild / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdS.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2c8.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhXW.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/559310a7-dbb0-461c-a863-5cb758607af5/f0474526-90d0-4d3d-aaae-dd68f3f38b28.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/75a9bdc4-3bc8-4be1-b77d-3774e6dde693/4bc73ecd-116a-422d-b80a-876586eb1e5d.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/a2bf8e68-b62f-413a-8b9d-b41b15f3b395/54061ddb-e264-4198-8f46-4a04fc4e2a3a.mp4"
        }
    },
    {
        "attribution": "Pro_Studio / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElb0.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe5U.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEab3.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/aa5cb260-7dae-44d3-acad-3c7053983ffe/1b790558-39a2-4d2a-bcd7-61f075e87fdd.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/346ec1aa-79a1-4da4-83c6-83474aabf815/cec6a12c-9775-4327-8ca3-28c6df4f8c74.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/a5418708-7037-447c-b54a-52cee3a36015/0552ee45-d456-4fb7-ad34-9f884d5b8515.mp4"
        }
    },
    {
        "attribution": "Leigh Prather / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEa3u.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEe5T.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElee.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/648cc494-7337-4644-9517-46877e93de76/486dfd9e-b104-4f2e-92a7-74c0eab6b14b.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/7d5a8818-8005-4647-afe0-c9955b578a7b/789ed99d-0fa5-4dc5-a4bd-942c44bc91b6.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/ebffe649-7bcc-49e1-a11e-dcfc96e2b1eb/705f3b7a-9c86-40cf-b2d5-65f2d24f9651.mp4"
        }
    },
    {
        "attribution": "Leigh Prather / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEhRL.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiFv.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElef.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/0f5c44c4-8d4a-4108-8890-355d55033565/689ed95c-c8a6-433d-89e6-3838a2f507b2.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/9c5d905e-ff07-45d9-a3c0-324841e6a4e8/93f3480b-e705-4152-8f8b-a5ce248fdd6a.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/fabad317-4a27-48d3-80e4-e3394b5e3697/190dae26-265e-4465-bc73-44cd120e3a04.mp4"
        }
    },
    {
        "attribution": "KASPORSKI VITALI / Shutterstock",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEiEo.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOElcx.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEab2.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/6ea6a8be-4bdb-498c-a34a-6ce0b3c5fe81/06096b63-be6b-4ccc-a7d5-d6b488be6974.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/569c8433-e759-4fba-910f-df231f2fbaee/8f927717-6ad6-4ac7-b7b6-b2d25fca40cd.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/5b6b6273-32c0-43af-85c0-a5d27662dc31/4939cb70-f95b-43c4-878f-5fe1256a456f.mp4"
        }
    },
    {
        "attribution": "kibrick / Getty Images",
        "firstFrame": {
            "i1080": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEcdN.img",
            "i1440": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOE2c4.img",
            "i2160": "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AAOEaaY.img"
        },
        "video": {
            "v1080": "https://prod-streaming-video-msn-com.akamaized.net/54cabe89-17bd-4116-8908-8501ebe48f6d/2e8ba8c7-a691-46d9-b5ae-d18050ce273a.mp4",
            "v1440": "https://prod-streaming-video-msn-com.akamaized.net/3998bbc2-3a27-4e28-ba73-b35ae062de3b/98456a38-81c2-4494-a457-fc5a89edb1fc.mp4",
            "v2160": "https://prod-streaming-video-msn-com.akamaized.net/5b8fbd42-4462-497c-81ed-5847e599f326/b245667f-6815-4952-ba69-7ba1389a27c4.mp4"
        }
    }
]
*/
