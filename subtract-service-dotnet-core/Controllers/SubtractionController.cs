using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

public class SubtractParams
{
    public int num1 { get; set; }
    public int num2 { get; set; }
}

public class DifferenceResponse
{
    public int difference { get; set; }
}


namespace SubtractServiceDotnetCore.Controllers
{
    [Route("subtract")]
    [ApiController]
    public class SubtractionController : ControllerBase
    {
        // POST subtract
        [HttpPost]
        public DifferenceResponse Post([FromBody] SubtractParams value)
        {
            DifferenceResponse response = new DifferenceResponse();
            response.difference = value.num1 - value.num2;
            return response;
        }
    }
}
