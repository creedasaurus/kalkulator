using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace DistCalc.Controllers
{
    [Route("api")]
    public class ValuesController : Controller
    {
        public class postData
        {
            public double num1;
            public double num2;
        }
        private static double subtract(double num1, double num2) => num1 - num2;

        [Route("subtract")]
        [HttpPost]
        public IActionResult Post(double num1, double num2)
        {
            return Ok(subtract(num1, num2));
        }
    }
}
